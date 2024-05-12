package product

import (
	"database/sql"
	localError "eniqlo/pkg/error"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type IProductRepository interface {
	FindBySku(sku string) (Product, *localError.GlobalError)
	CreateProduct(req CreateProductRequest) *localError.GlobalError
	FindAll(params QueryParams) ([]Product, *localError.GlobalError)
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) IProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) FindAll(params QueryParams) ([]Product, *localError.GlobalError) {
	products := []Product{}

	query := "SELECT * FROM products"
	nwhere := 0

	if params.ID != "" {
		nwhere += 1
		query += fmt.Sprintf(" WHERE id = '%s'", params.ID)
	}

	if params.Name != "" {
		prefix := "WHERE"
		if nwhere > 0 {
			prefix = "AND"
		}
		nwhere += 1
		query += fmt.Sprintf(" %s name ILIKE '%%%s%%'", prefix, params.Name)
	}

	if params.IsAvailable == "true" || params.IsAvailable == "false" {
		prefix := "WHERE"
		if nwhere > 0 {
			prefix = "AND"
		}
		nwhere += 1
		query += fmt.Sprintf(" %s is_available = %s", prefix, params.IsAvailable)
	}

	if params.Category == "Clothing" || params.Category == "Accessories" || params.Category == "Footwear" || params.Category == "Beverages" {
		prefix := "WHERE"
		if nwhere > 0 {
			prefix = "AND"
		}
		nwhere += 1
		query += fmt.Sprintf(" %s category = '%s'", prefix, params.Category)
	}

	if params.Sku != "" {
		prefix := "WHERE"
		if nwhere > 0 {
			prefix = "AND"
		}
		nwhere += 1
		query += fmt.Sprintf(" %s sku = '%s'", prefix, params.Sku)
	}

	if params.InStock == "true" {
		prefix := "WHERE"
		if nwhere > 0 {
			prefix = "AND"
		}
		nwhere += 1
		query += fmt.Sprintf(" %s stock > 0", prefix)
	} else if params.InStock == "false" {
		prefix := "WHERE"
		if nwhere > 0 {
			prefix = "AND"
		}
		nwhere += 1
		query += fmt.Sprintf(" %s stock = 0", prefix)
	}

	if params.Price == "asc" || params.Price == "desc" {
		query += fmt.Sprintf(" ORDER BY price %s", params.Price)
	} else if params.CreatedAt == "asc" || params.CreatedAt == "desc" {
		query += fmt.Sprintf(" ORDER BY created_at %s", params.CreatedAt)
	}

	if params.Limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", params.Limit)
	} else {
		query += " LIMIT 5"
	}
	if params.Offset != 0 {
		query += fmt.Sprintf(" OFFSET %d", params.Offset)
	} else {
		query += " OFFSET 0"
	}

	err := r.db.Select(&products, query)
	if err != nil {
		return products, localError.ErrInternalServer("Failed to find customers", err)
	}

	return products, nil
}

func (r *productRepository) FindBySku(sku string) (Product, *localError.GlobalError) {
	product := Product{}
	err := r.db.Get(&product, "SELECT * FROM products WHERE sku = $1", sku)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return product, localError.ErrInternalServer("Failed to find product", err)
		}
	}

	return product, nil
}

func (r *productRepository) CreateProduct(req CreateProductRequest) *localError.GlobalError {
	_, err := r.db.Exec("INSERT INTO products (name, sku, category, image_url, notes, price, stock, location, is_available) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		req.Name, req.Sku, req.Category, req.ImageURL, req.Notes, req.Price, req.Stock, req.Location, req.IsAvailable)
	if err != nil {
		return localError.ErrInternalServer("Failed to create product", err)
	}

	return nil
}

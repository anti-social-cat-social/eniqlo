package product

import (
	"database/sql"
	localError "eniqlo/pkg/error"
	"errors"

	"github.com/jmoiron/sqlx"
)

type IProductRepository interface {
	FindBySku(sku string) (Product, *localError.GlobalError)
	CreateProduct(req CreateProductRequest) *localError.GlobalError
	FindByID(id string) (Product, *localError.GlobalError)
	DeleteProduct(id string) *localError.GlobalError
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) IProductRepository {
	return &productRepository{
		db: db,
	}
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

func (r *productRepository) FindByID(id string) (Product, *localError.GlobalError) {
	product := Product{}
	err := r.db.Get(&product, "SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return product, localError.ErrInternalServer("Failed to find product", err)
		}
	}

	return product, nil
}

func (r *productRepository) DeleteProduct(id string) *localError.GlobalError {
	_, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return localError.ErrInternalServer("Failed to delete product", err)
	}

	return nil
}

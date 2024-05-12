package customer

import (
	localError "eniqlo/pkg/error"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ICustomerRepository interface {
	FindAll(params QueryParams) ([]Customer, *localError.GlobalError)
}

type customerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) ICustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) FindAll(params QueryParams) ([]Customer, *localError.GlobalError) {
	customers := []Customer{}

	query := "SELECT * FROM customers"
	if params.Name != "" {
		query += fmt.Sprintf(" WHERE name ILIKE '%%%s%%'", params.Name)
	}
	if params.PhoneNumber != "" {
		if params.Name != "" {
			query += fmt.Sprintf(" AND phone_number ILIKE '%%%s%%'", params.PhoneNumber)
		} else {
			query += fmt.Sprintf(" WHERE phone_number ILIKE '%%%s%%'", params.PhoneNumber)
		}
	}
	if params.Limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", params.Limit)
	} else {
		query += " LIMIT 10"
	}
	if params.Offset != 0 {
		query += fmt.Sprintf(" OFFSET %d", params.Offset)
	} else {
		query += " OFFSET 0"
	}

	fmt.Println(query)
	err := r.db.Select(&customers, query)
	if err != nil {
		return customers, localError.ErrInternalServer("Failed to find customers", err)
	}

	return customers, nil
}

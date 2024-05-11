package customer

import (
	localError "eniqlo/pkg/error"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type ICustomerRepository interface {
	FindByPhone(phone string) (*Customer, *localError.GlobalError)
	Create(entity Customer) (*Customer, *localError.GlobalError)
}

type customerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) ICustomerRepository {
	return &customerRepository{
		db: db,
	}
}

// Store new customer to database
func (u *customerRepository) Create(entity Customer) (*Customer, *localError.GlobalError) {
	// entity.ID = uuid.NewString()

	// Insert into database
	query := "INSERT INTO customers (id, phone_number, name) values (:id, :phone_number, :name);"
	_, err := u.db.NamedExec(query, &entity)
	if err != nil {
		return nil, localError.ErrInternalServer(err.Error(), err)
	}

	return &entity, nil
}

// Find customer by phone
// This can be use for authentication process
func (u *customerRepository) FindByPhone(phone string) (*Customer, *localError.GlobalError) {
	var customer Customer

	if err := u.db.Get(&customer, "SELECT * FROM customers where phone_number=$1", phone); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, localError.ErrNotFound("Customer data not found", err)
		}

		return nil, &localError.GlobalError{
			Code:    400,
			Message: "Not found",
			Error:   err,
		}

	}

	return &customer, nil
}

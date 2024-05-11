package customer

import (
	localError "eniqlo/pkg/error"
	"errors"

	"github.com/google/uuid"
)

type ICustomerUsecase interface {
	Register(dto CustomerRegisterDTO) (*CustomerRegisterResponse, *localError.GlobalError)
}

type customerUsecase struct {
	repo ICustomerRepository
}

func NewCustomerUsecase(repo ICustomerRepository) ICustomerUsecase {
	return &customerUsecase{
		repo: repo,
	}
}

// Register implements ICustomerUsecase.
func (u *customerUsecase) Register(dto CustomerRegisterDTO) (*CustomerRegisterResponse, *localError.GlobalError) {
	// Validate customer request first

	// Check if customer with given phone is already exists
	existedCustomer, _ := u.repo.FindByPhone(dto.PhoneNumber)
	if existedCustomer != nil {
		return nil, localError.ErrConflict("Customer already exists", errors.New("customer already exists"))
	}

	// Map DTO to customer entity
	// This used for storing data to database
	customer := Customer{
		ID: uuid.NewString(),
		Name:  dto.Name,
		PhoneNumber: dto.PhoneNumber,
	}

	registeredCustomer, err := u.repo.Create(customer)
	if err != nil {
		return nil, err
	}

	response := CustomerRegisterResponse{
		UserId: registeredCustomer.ID,
		PhoneNumber: registeredCustomer.PhoneNumber, 
		Name: registeredCustomer.Name,
	}

	return &response, nil
}
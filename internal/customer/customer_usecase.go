package customer

import (
	localError "eniqlo/pkg/error"
)

type ICustomerUsecase interface {
	FindCustomers(query QueryParams) ([]Customer, *localError.GlobalError)
}

type customerUsecase struct {
	repo ICustomerRepository
}

func NewCustomerUsecase(repo ICustomerRepository) ICustomerUsecase {
	return &customerUsecase{
		repo: repo,
	}
}

func (uc *customerUsecase) FindCustomers(query QueryParams) ([]Customer, *localError.GlobalError) {
	customers, err := uc.repo.FindAll(query)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

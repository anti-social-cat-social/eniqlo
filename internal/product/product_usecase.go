package product

import (
	localError "eniqlo/pkg/error"
	"eniqlo/pkg/validation"
)

type IProductUsecase interface {
	CreateProduct(req CreateProductRequest) (*Product, *localError.GlobalError)
}

type productUsecase struct {
	repo IProductRepository
}

func NewProductUsecase(repo IProductRepository) IProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}

func (uc *productUsecase) CreateProduct(req CreateProductRequest) (*Product, *localError.GlobalError) {
	if !validation.IsValidURL(req.ImageURL) {
		return nil, localError.ErrBadRequest("Invalid image URL", nil)
	}

	product, err := uc.repo.FindBySku(req.Sku)
	if err != nil {
		return nil, err
	}

	if product.ID != "" {
		return nil, localError.ErrBadRequest("SKU already exists", nil)
	}

	err = uc.repo.CreateProduct(req)
	if err != nil {
		return nil, err
	}

	product, err = uc.repo.FindBySku(req.Sku)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

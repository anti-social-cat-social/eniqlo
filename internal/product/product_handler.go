package product

import (
	"eniqlo/internal/middleware"
	"eniqlo/pkg/response"
	"eniqlo/pkg/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	uc IProductUsecase
}

func NewProductHandler(uc IProductUsecase) *productHandler {
	return &productHandler{
		uc: uc,
	}
}

func (h *productHandler) Router(r *gin.RouterGroup) {
	group := r.Group("product")

	group.Use(middleware.UseJwtAuth)

	group.POST("", h.CreateProduct)
	group.GET("", h.FindAll)
}

func (h *productHandler) FindAll(c *gin.Context) {
	query := QueryParams{}
	if err := c.ShouldBindQuery(&query); err != nil {
		res := validation.FormatValidation(err)
		response.GenerateResponse(c, res.Code, response.WithMessage(res.Message))
		return
	}

	products, err := h.uc.FindProducts(query)
	if err != nil {
		response.GenerateResponse(c, err.Code, response.WithMessage(err.Message))
		return
	}

	// res := FormatCustomersResponse(products)

	response.GenerateResponse(c, http.StatusOK, response.WithMessage("Product fetched successfully!"), response.WithData(products))
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var request CreateProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		res := validation.FormatValidation(err)
		response.GenerateResponse(c, res.Code, response.WithMessage(res.Message))
		return
	}

	product, err := h.uc.CreateProduct(request)
	if err != nil {
		response.GenerateResponse(c, err.Code, response.WithMessage(err.Message))
		return
	}

	res := FormatCreateProductResponse(*product)

	response.GenerateResponse(c, http.StatusCreated, response.WithMessage("Product created successfully!"), response.WithData(res))
}

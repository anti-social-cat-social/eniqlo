package customer

import (
	"eniqlo/internal/middleware"
	"eniqlo/pkg/response"
	"eniqlo/pkg/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	uc ICustomerUsecase
}

func NewCustomerHandler(uc ICustomerUsecase) *customerHandler {
	return &customerHandler{
		uc: uc,
	}
}

func (h *customerHandler) Router(r *gin.RouterGroup) {
	group := r.Group("customer")

	group.Use(middleware.UseJwtAuth)

	group.GET("", h.FindAll)
}

func (h *customerHandler) FindAll(c *gin.Context) {
	query := QueryParams{}
	if err := c.ShouldBindQuery(&query); err != nil {
		res := validation.FormatValidation(err)
		response.GenerateResponse(c, res.Code, response.WithMessage(res.Message))
		return
	}

	customers, err := h.uc.FindCustomers(query)
	if err != nil {
		response.GenerateResponse(c, err.Code, response.WithMessage(err.Message))
		return
	}

	response.GenerateResponse(c, http.StatusOK, response.WithMessage("Customer fetched successfully!"), response.WithData(customers))
}

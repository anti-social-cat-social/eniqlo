package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRoute(engine *gin.Engine, db *sqlx.DB) {
	// Handle for not found routes
	router := engine.Group("v1")

	router.GET("ping", pingHandler)
}

func pingHandler(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		struct {
			Data    any    `json:"data"`
			Message string `json:"message"`
			Success bool   `json:"success"`
		}{
			Success: true,
			Message: "Server is online",
			Data:    true,
		},
	)
}

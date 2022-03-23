package handler

import (
	"github.com/gin-gonic/gin"
	"historics-api/internal/service"
)

type Handler struct {
	HistoricService    service.HistoricService
	TransactionService service.TransactionService
}

type Config struct {
	Router             *gin.Engine
	HistoricService    service.HistoricService
	TransactionService service.TransactionService
}

func NewHandler(c *Config) {
	h := &Handler{
		HistoricService:    c.HistoricService,
		TransactionService: c.TransactionService,
	}

	api := c.Router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/", h.Health)
			v1.GET("/:collection/historic", h.GetHistoric)
			v1.GET("/:collection/transactions", h.GetTransactions)
		}
	}
}

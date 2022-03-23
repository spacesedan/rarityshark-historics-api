package internal

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"historics-api/internal/handler"
	"historics-api/internal/repo"
	"historics-api/internal/service"
)

func Inject(db *mongo.Database) (*gin.Engine, error) {
	dao := repo.NewDAO(db)
	historic := service.NewHistoricService(dao)
	transaction := service.NewTransactionService(dao)

	app := gin.Default()

	handler.NewHandler(&handler.Config{
		Router:             app,
		HistoricService:    historic,
		TransactionService: transaction,
	})

	return app, nil
}

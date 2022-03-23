package repo

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

var (
	metaCollection         = "META-collections"
	historicCollection     = "-historic"
	transactionsCollection = "-transactions"
)

type DAO interface {
	NewHistoricQuery() HistoricsQuery
	NewTransactionQuery() TransactionQuery
}

type dao struct{}

var DB *mongo.Database

func NewDAO(db *mongo.Database) DAO {
	DB = db
	return &dao{}
}

func NewMongo() (*mongo.Database, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("Failed to connect to mongo")
		return nil, err
	}

	fmt.Println("Connected to mongodb")

	db := client.Database(os.Getenv("DB"))

	return db, nil
}

func (d *dao) NewHistoricQuery() HistoricsQuery {
	return &historicsQuery{}
}

func (d *dao) NewTransactionQuery() TransactionQuery {
	return &transactionQuery{}
}

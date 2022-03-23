package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"historics-api/internal/models"
	"time"
)

type TransactionQuery interface {
	GetTransactions(cn string, limit, page int64) ([]*models.TokenTransaction, error)
}

type transactionQuery struct {
}

func (t *transactionQuery) GetTransactions(cn string, limit, page int64) ([]*models.TokenTransaction, error) {
	m := DB.Collection(cn + transactionsCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := bson.D{{}}

	opts := options.Find()
	opts.SetLimit(limit)
	opts.SetSkip(page)
	opts.SetSort(bson.D{{"time", -1}})

	var data []*models.TokenTransaction

	res, err := m.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer res.Close(ctx)
	for res.Next(ctx) {
		var doc *models.TokenTransaction
		res.Decode(&doc)
		data = append(data, doc)
	}

	return data, nil
}

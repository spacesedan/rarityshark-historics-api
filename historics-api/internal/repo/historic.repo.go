package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"historics-api/internal/models"
	"time"
)

type HistoricsQuery interface {
	GetHistorics(cn string, limit int64) ([]*models.Historic, error)
}

type historicsQuery struct {
}

func (h *historicsQuery) GetHistorics(cn string, limit int64) ([]*models.Historic, error) {
	m := DB.Collection(cn + historicCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data []*models.Historic

	opts := options.Find()
	opts.SetLimit(limit)

	results, err := m.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func(results *mongo.Cursor, ctx context.Context) {
		err := results.Close(ctx)
		if err != nil {

		}
	}(results, ctx)
	for results.Next(ctx) {
		var doc *models.Historic
		err := results.Decode(&doc)
		if err != nil {
			return nil, err
		}
		data = append(data, doc)
	}

	return data, nil

}

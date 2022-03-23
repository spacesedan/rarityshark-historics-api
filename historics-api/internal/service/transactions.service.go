package service

import (
	"historics-api/internal/dto"
	"historics-api/internal/models"
	"historics-api/internal/repo"
	"log"
)

type TransactionService interface {
	GetTransactions(dto dto.GetTransaction) ([]*models.TokenTransaction, error)
}

type transactionService struct {
	dao repo.DAO
}

func NewTransactionService(dao repo.DAO) TransactionService {
	return &transactionService{
		dao: dao,
	}
}

func (t *transactionService) GetTransactions(dto dto.GetTransaction) ([]*models.TokenTransaction, error) {
	txs, err := t.dao.NewTransactionQuery().GetTransactions(dto.Collection, dto.Limit, dto.Page)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return txs, nil
}

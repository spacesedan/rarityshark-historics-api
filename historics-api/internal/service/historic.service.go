package service

import (
	"historics-api/internal/dto"
	"historics-api/internal/models"
	"historics-api/internal/repo"
)

type HistoricService interface {
	GetHistoricData(dto dto.GetHistorics) ([]*models.Historic, error)
}

type historicService struct {
	dao repo.DAO
}

func NewHistoricService(dao repo.DAO) HistoricService {
	return &historicService{
		dao: dao,
	}
}

func (h *historicService) GetHistoricData(dto dto.GetHistorics) ([]*models.Historic, error) {
	data, err := h.dao.NewHistoricQuery().GetHistorics(dto.Collection, dto.Limit)
	if err != nil {
		return nil, err
	}

	return data, nil
}

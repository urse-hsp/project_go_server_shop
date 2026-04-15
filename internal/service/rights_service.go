package service

import (
	"context"
	"go-server/internal/dao"
	"go-server/internal/model"
)

type RightsService interface {
	GetList(ctx context.Context) ([]model.Permission, error)
}

func NewRightsService(
	service *Service,
	Repo dao.RightsRepository,
) RightsService {
	return &rightsService{
		Repo:    Repo,
		Service: service,
	}
}

type rightsService struct {
	*Service
	Repo dao.RightsRepository
}

// ================= 全部列表 =================

func (s *rightsService) GetList(ctx context.Context) ([]model.Permission, error) {
	return s.Repo.GetList(ctx)
}

package service

import (
	"context"
	"fmt"
	"go-server/internal/dao"
	"go-server/internal/model"
	"go-server/pkg/bcrypt"
	"time"
)

// manager 管理员
type ManagerService interface {
	Login(ctx context.Context, username string, password string) (*model.Manager, string, error)
	GetManagerLists(ctx context.Context, page, pageSize int) ([]model.Manager, int64, error)
}

type managerService struct {
	*Service
	managerRepo dao.ManagerRepository
}

func NewManagerService(
	service *Service,
	managerRepo dao.ManagerRepository,
) ManagerService {
	return &managerService{
		managerRepo: managerRepo,
		Service:     service,
	}
}

func (s *managerService) Login(ctx context.Context, username string, password string) (*model.Manager, string, error) {
	user, err := s.managerRepo.GetManagerByUsername(ctx, username)
	if err != nil {
		return nil, "", err
	}

	if !bcrypt.CheckPassword(password, user.MgPwd) {
		return nil, "", fmt.Errorf("密码错误")
	}

	// token, err := jwt.GenerateToken(user.ID, user.Username)
	// duration := time.Duration(s.) * time.Hour
	token, err := s.jwt.GenToken(user.MgID, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *managerService) GetManagerLists(ctx context.Context, page, pageSize int) ([]model.Manager, int64, error) {
	return s.managerRepo.GetManagerLists(ctx, page, pageSize)
}

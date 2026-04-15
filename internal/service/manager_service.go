package service

import (
	"context"
	"errors"
	"fmt"
	"go-server/internal/dao"
	managerdto "go-server/internal/dto/manager"
	"go-server/internal/model"
	"go-server/pkg/bcrypt"
	"time"

	"gorm.io/gorm"
)

// manager 管理员
type ManagerService interface {
	Login(ctx context.Context, username string, password string) (*model.Manager, string, error)

	Create(ctx context.Context, req managerdto.CreateRequest) (*model.Manager, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, req managerdto.UpdateRequest) (*model.Manager, error)
	// GetDetail(ctx context.Context, id uint) (*model.Manager, error)
	GetLists(ctx context.Context, q managerdto.ManagerQuery) ([]model.Manager, int64, error)
}

type managerService struct {
	*Service
	Repo dao.ManagerRepository
}

func NewManagerService(
	service *Service,
	Repo dao.ManagerRepository,
) ManagerService {
	return &managerService{
		Repo:    Repo,
		Service: service,
	}
}

func (s *managerService) Login(ctx context.Context, username string, password string) (*model.Manager, string, error) {
	data, err := s.Repo.GetByKeyWhere(ctx, username)
	if err != nil {
		return nil, "", err
	}

	if !bcrypt.CheckPassword(password, data.MgPwd) {
		return nil, "", fmt.Errorf("密码错误")
	}

	// token, err := jwt.GenerateToken(data.ID, data.Username)
	// duration := time.Duration(s.) * time.Hour
	token, err := s.jwt.GenToken(data.MgID, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return nil, "", err
	}

	return data, token, nil
}

// ================= 注册 =================

func (s *managerService) Create(ctx context.Context, req managerdto.CreateRequest) (*model.Manager, error) {
	hashedPwd, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	data := &model.Manager{
		MgName:   req.Username,
		MgPwd:    hashedPwd,
		MgMobile: req.Mobile,
		MgEmail:  req.Email,
	}
	return s.Repo.Create(ctx, data)
}

func (s *managerService) GetLists(ctx context.Context, q managerdto.ManagerQuery) ([]model.Manager, int64, error) {
	return s.Repo.GetLists(ctx, q)
}

// ================= 删除 =================

func (s *managerService) Delete(ctx context.Context, id uint) error {
	// 直接调用 DAO（DAO 已经判断是否存在）
	return s.Repo.Delete(ctx, id)
}

// ================= 更新 =================

func (s *managerService) Update(ctx context.Context, id uint, req managerdto.UpdateRequest) (*model.Manager, error) {
	_, err := s.Repo.GetDetail(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户不存在")
		}
		return nil, err
	}

	data := &model.Manager{}

	AssignIfNotNil(&data.MgMobile, req.Mobile)
	AssignIfNotNil(&data.MgEmail, req.Email)

	if req.State != nil {
		if *req.State {
			data.MgState = 1
		} else {
			data.MgState = 0
		}
	}

	return s.Repo.Update(ctx, data, id)
}

package service

import (
	"context"
	"errors"
	"fmt"
	"go-server/internal/dao"
	roledto "go-server/internal/dto/role"
	"go-server/internal/model"

	"gorm.io/gorm"
)

type RoleService interface {
	Create(ctx context.Context, req roledto.LoginRequest) (*model.Role, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, req roledto.LoginRequest) (*model.Role, error)
	GetList(ctx context.Context) ([]model.Role, error)
	GetLists(ctx context.Context, page, pageSize int) ([]model.Role, int64, error)
	GetDetail(ctx context.Context, id uint) (*model.Role, error)
}

func NewRoleService(
	service *Service,
	roleRepo dao.RoleRepository,
) RoleService {
	return &roleService{
		roleRepo: roleRepo,
		Service:  service,
	}
}

type roleService struct {
	*Service
	roleRepo dao.RoleRepository
}

// ================= 注册 =================

func (s *roleService) Create(ctx context.Context, req roledto.LoginRequest) (*model.Role, error) {
	// 判断是否已存在
	role, err := s.roleRepo.GetByKeyWhere(ctx, req.RoleName)
	if err == nil && role != nil {
		return nil, fmt.Errorf("名称已存在")
	}

	data := &model.Role{
		RoleName: req.RoleName,
		RoleDesc: req.RoleDesc,
	}

	return s.roleRepo.Create(ctx, data)
}

// ================= 删除 =================

func (s *roleService) Delete(ctx context.Context, id uint) error {
	return s.roleRepo.Delete(ctx, id)
}

// ================= 更新 =================

func (s *roleService) Update(ctx context.Context, id uint, req roledto.LoginRequest) (*model.Role, error) {
	role, err := s.roleRepo.GetDetail(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户不存在")
		}
		return nil, err
	}

	role.RoleName = req.RoleName
	role.RoleDesc = req.RoleDesc

	return s.roleRepo.Update(ctx, role, id)
}

// ================= 获取 =================

func (s *roleService) GetDetail(ctx context.Context, id uint) (*model.Role, error) {
	return s.roleRepo.GetDetail(ctx, id)
}

// ================= 全部列表 =================

func (s *roleService) GetList(ctx context.Context) ([]model.Role, error) {
	return s.roleRepo.GetList(ctx)
}

// ================= 分页列表 =================

func (s *roleService) GetLists(ctx context.Context, page, pageSize int) ([]model.Role, int64, error) {
	return s.roleRepo.GetLists(ctx, page, pageSize)
}

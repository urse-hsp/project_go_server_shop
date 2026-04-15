package service

import (
	"context"
	"errors"
	"fmt"
	"go-server/internal/dao"
	category "go-server/internal/dto/category"
	"go-server/internal/model"

	"gorm.io/gorm"
)

type CategoryService interface {
	Create(ctx context.Context, req category.CreateRequest) (*model.Category, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, req category.UpdateRequest) (*model.Category, error)
	GetDetail(ctx context.Context, id uint) (*model.Category, error)
	GetList(ctx context.Context, q category.RequestQuery) ([]model.Category, error)
	GetPageList(ctx context.Context, q category.RequestPageQuery) ([]model.Category, int64, error)
}

func NewCategoryCatsService(
	service *Service,
	Repo dao.CategoryRepository,
) CategoryService {
	return &categoryService{
		Repo:    Repo,
		Service: service,
	}
}

type categoryService struct {
	*Service
	Repo dao.CategoryRepository
}

// ================= 创建 =================

func (s *categoryService) Create(ctx context.Context, req category.CreateRequest) (*model.Category, error) {
	// 判断是否已存在
	// data, err := s.Repo.GetByKeyWhere(ctx, req.CatName)
	// if err == nil && data != nil {
	// 	return nil, fmt.Errorf("数据已存在")
	// }

	data_ := &model.Category{
		CatName:  req.CatName,
		CatPID:   req.CatPID,
		CatLevel: req.CatLevel,
	}

	return s.Repo.Create(ctx, data_)
}

// ================= 删除 =================

func (s *categoryService) Delete(ctx context.Context, id uint) error {
	// 直接调用 DAO（DAO 已经判断是否存在）
	return s.Repo.Delete(ctx, id)
}

// ================= 更新 =================

func (s *categoryService) Update(ctx context.Context, id uint, req category.UpdateRequest) (*model.Category, error) {
	data, err := s.Repo.GetDetail(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("数据不存在")
		}
		return nil, err
	}

	// AssignIfNotNil(&data.CatName, req.CatName)
	data.CatName = req.CatName

	return s.Repo.Update(ctx, data, id)
}

// ================= 获取 =================

func (s *categoryService) GetDetail(ctx context.Context, id uint) (*model.Category, error) {
	return s.Repo.GetDetail(ctx, id)
}

// ================= 全部列表 =================

func (s *categoryService) GetList(ctx context.Context, q category.RequestQuery) ([]model.Category, error) {
	return s.Repo.GetList(ctx, q)
}

// ================= 分页列表 =================

func (s *categoryService) GetPageList(ctx context.Context, q category.RequestPageQuery) ([]model.Category, int64, error) {
	return s.Repo.GetPageList(ctx, q)
}

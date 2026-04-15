package service

import (
	"context"
	"errors"
	"fmt"
	"go-server/internal/dao"
	attributedto "go-server/internal/dto/attribute"
	"go-server/internal/model"

	"gorm.io/gorm"
)

type AttributeService interface {
	Create(ctx context.Context, id uint, req attributedto.CreateRequest) (*model.Attribute, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, req attributedto.UpdateRequest) (*model.Attribute, error)
	GetDetail(ctx context.Context, id uint) (*model.Attribute, error)
	GetList(ctx context.Context, id uint, q attributedto.RequestQuery) ([]model.Attribute, error)
	GetPageList(ctx context.Context, id uint, q attributedto.RequestPageQuery) ([]model.Attribute, int64, error)
}

func NewAttributeService(
	service *Service,
	Repo dao.AttributeRepository,
) AttributeService {
	return &attributeService{
		Repo:    Repo,
		Service: service,
	}
}

type attributeService struct {
	*Service
	Repo dao.AttributeRepository
}

// ================= 创建 =================

func (s *attributeService) Create(ctx context.Context, id uint, req attributedto.CreateRequest) (*model.Attribute, error) {
	// 判断是否已存在
	data, err := s.Repo.GetByKeyWhere(ctx, id, req.AttrName)
	if err == nil && data != nil {
		return nil, fmt.Errorf("数据已存在")
	}

	data_ := &model.Attribute{
		AttrName: req.AttrName,
		CatID:    id,
		AttrSel:  req.Sel,
	}

	return s.Repo.Create(ctx, data_)
}

// ================= 删除 =================

func (s *attributeService) Delete(ctx context.Context, id uint) error {
	// 直接调用 DAO（DAO 已经判断是否存在）
	return s.Repo.Delete(ctx, id)
}

// ================= 更新 =================

func (s *attributeService) Update(ctx context.Context, id uint, req attributedto.UpdateRequest) (*model.Attribute, error) {
	data, err := s.Repo.GetDetail(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("数据不存在")
		}
		return nil, err
	}

	data.AttrName = req.AttrName
	AssignIfNotNil(&data.AttrVals, req.AttrVals)

	return s.Repo.Update(ctx, data, id)
}

// ================= 获取 =================

func (s *attributeService) GetDetail(ctx context.Context, id uint) (*model.Attribute, error) {
	return s.Repo.GetDetail(ctx, id)
}

// ================= 全部列表 =================

func (s *attributeService) GetList(ctx context.Context, id uint, q attributedto.RequestQuery) ([]model.Attribute, error) {
	return s.Repo.GetList(ctx, id, q)
}

// ================= 分页列表 =================

func (s *attributeService) GetPageList(ctx context.Context, id uint, q attributedto.RequestPageQuery) ([]model.Attribute, int64, error) {
	return s.Repo.GetPageList(ctx, id, q)
}

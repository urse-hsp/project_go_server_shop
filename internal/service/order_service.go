package service

import (
	"context"
	"errors"
	"fmt"
	"go-server/internal/dao"
	orderdto "go-server/internal/dto/order"
	"go-server/internal/model"

	"gorm.io/gorm"
)

type OrderService interface {
	Create(ctx context.Context, req orderdto.CreateRequest) (*model.Order, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, req orderdto.UpdateRequest) (*model.Order, error)
	GetDetail(ctx context.Context, id uint) (*model.Order, error)
	GetList(ctx context.Context, q orderdto.RequestQuery) ([]model.Order, error)
	GetPageList(ctx context.Context, q orderdto.RequestPageQuery) ([]model.Order, int64, error)
}

func NewOrderService(
	service *Service,
	Repo dao.OrderRepository,
) OrderService {
	return &orderService{
		Repo:    Repo,
		Service: service,
	}
}

type orderService struct {
	*Service
	Repo dao.OrderRepository
}

// ================= 创建 =================

func (s *orderService) Create(ctx context.Context, req orderdto.CreateRequest) (*model.Order, error) {

	data_ := &model.Order{}

	return s.Repo.Create(ctx, data_)
}

// ================= 删除 =================

func (s *orderService) Delete(ctx context.Context, id uint) error {
	// 直接调用 DAO（DAO 已经判断是否存在）
	return s.Repo.Delete(ctx, id)
}

// ================= 更新 =================

func (s *orderService) Update(ctx context.Context, id uint, req orderdto.UpdateRequest) (*model.Order, error) {
	data, err := s.Repo.GetDetail(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("数据不存在")
		}
		return nil, err
	}

	// ✅ 处理 IsSend
	if req.IsSend != nil {
		switch *req.IsSend {
		case orderdto.IsSendYes:
			data.IsSend = "是"
		case orderdto.IsSendNo:
			data.IsSend = "否"
		default:
			return nil, fmt.Errorf("is_send 参数非法")
		}
	}

	AssignIfNotNil(&data.OrderPay, req.OrderPay)
	AssignIfNotNil(&data.OrderPrice, req.OrderPrice)
	AssignIfNotNil(&data.PayStatus, req.PayStatus)

	return s.Repo.Update(ctx, data, id)
}

// ================= 获取 =================

func (s *orderService) GetDetail(ctx context.Context, id uint) (*model.Order, error) {
	return s.Repo.GetDetail(ctx, id)
}

// ================= 全部列表 =================

func (s *orderService) GetList(ctx context.Context, q orderdto.RequestQuery) ([]model.Order, error) {
	return s.Repo.GetList(ctx, q)
}

// ================= 分页列表 =================

func (s *orderService) GetPageList(ctx context.Context, q orderdto.RequestPageQuery) ([]model.Order, int64, error) {
	return s.Repo.GetPageList(ctx, q)
}

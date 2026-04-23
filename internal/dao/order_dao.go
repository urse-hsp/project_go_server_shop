package dao

import (
	"context"
	"fmt"
	orderdto "go-server/internal/dto/order"
	"go-server/internal/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, data *model.Order) (*model.Order, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, data *model.Order, id uint) (*model.Order, error)
	GetDetail(ctx context.Context, id uint) (*model.Order, error)
	GetList(ctx context.Context, q orderdto.RequestQuery) ([]model.Order, error)
	GetPageList(ctx context.Context, q orderdto.RequestPageQuery) ([]model.Order, int64, error)

	buildQuery(ctx context.Context, q orderdto.RequestQuery) *gorm.DB
	GetByKeyWhere(ctx context.Context, username string) (*model.Order, error)
}

func NewOrderRepository(
	r *Repository,
) OrderRepository {
	return &orderRepository{
		Repository: r,
	}
}

type orderRepository struct {
	*Repository
}

// ================= 根据ID查询 =================

func (r *orderRepository) GetDetail(ctx context.Context, id uint) (*model.Order, error) {
	var data model.Order

	err := r.DB(ctx).First(&data, id).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ================= 根据关键字查询 =================

func (r *orderRepository) GetByKeyWhere(ctx context.Context, username string) (*model.Order, error) {
	var data model.Order

	err := r.DB(ctx).Where("order_number = ?", username).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ================= 创建 =================

func (r *orderRepository) Create(ctx context.Context, data *model.Order) (*model.Order, error) {

	if err := r.DB(ctx).Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// ================= 更新 =================

func (r *orderRepository) Update(ctx context.Context, data *model.Order, id uint) (*model.Order, error) {
	if err := r.DB(ctx).
		Model(&model.Order{}).
		Where("order_id = ?", id).
		Updates(data).Error; err != nil {
		return nil, err
	}

	return r.GetDetail(ctx, id)
}

// ================= 删除 =================

func (r *orderRepository) Delete(ctx context.Context, id uint) error {
	result := r.DB(ctx).Where("order_id = ?", id).Delete(&model.Order{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("id不存在")
	}

	return nil
}

// ================= 全部列表 =================

func (r *orderRepository) GetList(ctx context.Context, q orderdto.RequestQuery) ([]model.Order, error) {
	var data []model.Order

	db := r.buildQuery(ctx, q)

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// ================= 分页列表 =================

func (r *orderRepository) GetPageList(ctx context.Context, q orderdto.RequestPageQuery) ([]model.Order, int64, error) {
	var data []model.Order

	db := r.buildQuery(ctx, q.RequestQuery)

	total, err := Paginate(db, &data, q.Page, q.PageSize)

	return data, total, err
}

// ================= 公共查询 =================
func (r *orderRepository) buildQuery(ctx context.Context, q orderdto.RequestQuery) *gorm.DB {
	db := r.DB(ctx).Model(&model.Order{})

	if q.Query != nil {
		like := "%" + *q.Query + "%"
		db = db.Where("order_number LIKE ?", like)
	}

	return db
}

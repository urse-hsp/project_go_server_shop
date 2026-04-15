package dao

import (
	"context"
	"fmt"
	v1 "go-server/api/v1"
	"go-server/internal/bootstrap"
	"go-server/internal/model"
)

type RightsRepository interface {
	GetByUsername(ctx context.Context, username string) (*model.Permission, error)

	Create(ctx context.Context, user *model.Permission) (*model.Permission, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, user *model.Permission, id uint) (*model.Permission, error)
	GetDetail(ctx context.Context, id uint) (*model.Permission, error)
	GetList(ctx context.Context) ([]model.Permission, error)
	GetLists(ctx context.Context, q v1.PageRequest) ([]model.Permission, int64, error)
}

func NewRightsRepository(
	r *bootstrap.Repository,
) RightsRepository {
	return &rightsRepository{
		Repository: r,
	}
}

type rightsRepository struct {
	*bootstrap.Repository
}

// ================= 根据ID查询 =================

func (r *rightsRepository) GetDetail(ctx context.Context, id uint) (*model.Permission, error) {
	var data model.Permission

	err := r.DB(ctx).First(&data, id).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ================= 根据关键字查询 =================

func (r *rightsRepository) GetByUsername(ctx context.Context, username string) (*model.Permission, error) {
	var data model.Permission

	err := r.DB(ctx).Where("username = ?", username).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ================= 创建 =================

func (r *rightsRepository) Create(ctx context.Context, data *model.Permission) (*model.Permission, error) {

	if err := r.DB(ctx).Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// ================= 更新 =================

func (r *rightsRepository) Update(ctx context.Context, data *model.Permission, id uint) (*model.Permission, error) {
	if err := r.DB(ctx).
		Model(&model.Permission{}).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return nil, err
	}

	return r.GetDetail(ctx, id)
}

// ================= 删除 =================

func (r *rightsRepository) Delete(ctx context.Context, id uint) error {
	result := r.DB(ctx).Where("id = ?", id).Delete(&model.Permission{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("用户不存在")
	}

	return nil
}

// ================= 全部列表 =================

func (r *rightsRepository) GetList(ctx context.Context) ([]model.Permission, error) {
	var data []model.Permission

	if err := r.DB(ctx).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// ================= 分页列表 =================

func (r *rightsRepository) GetLists(ctx context.Context, q v1.PageRequest) ([]model.Permission, int64, error) {
	var data []model.Permission

	db := r.DB(ctx).Model(&model.Permission{})

	// 分页
	total, err := Paginate(db, &data, q.Page, q.PageSize)

	return data, total, err
}

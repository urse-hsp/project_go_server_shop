package dao

import (
	"context"
	"fmt"
	attributedto "go-server/internal/dto/attribute"
	"go-server/internal/model"

	"gorm.io/gorm"
)

type AttributeRepository interface {
	Create(ctx context.Context, data *model.Attribute) (*model.Attribute, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, data *model.Attribute, id uint) (*model.Attribute, error)
	GetDetail(ctx context.Context, id uint) (*model.Attribute, error)
	GetList(ctx context.Context, id uint, q attributedto.RequestQuery) ([]model.Attribute, error)
	GetPageList(ctx context.Context, id uint, q attributedto.RequestPageQuery) ([]model.Attribute, int64, error)

	buildQuery(ctx context.Context, id uint, q attributedto.RequestQuery) *gorm.DB
	GetByKeyWhere(ctx context.Context, id uint, name string) (*model.Attribute, error)
}

func NewAttributeRepository(
	r *Repository,
) AttributeRepository {
	return &attributeRepository{
		Repository: r,
	}
}

type attributeRepository struct {
	*Repository
}

// ================= 根据ID查询 =================

func (r *attributeRepository) GetDetail(ctx context.Context, id uint) (*model.Attribute, error) {
	var data model.Attribute

	err := r.DB(ctx).First(&data, id).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ================= 根据关键字查询 =================

func (r *attributeRepository) GetByKeyWhere(ctx context.Context, id uint, name string) (*model.Attribute, error) {
	var data model.Attribute

	err := r.DB(ctx).Where("username = ?", name).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ================= 创建 =================

func (r *attributeRepository) Create(ctx context.Context, data *model.Attribute) (*model.Attribute, error) {

	if err := r.DB(ctx).Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// ================= 更新 =================

func (r *attributeRepository) Update(ctx context.Context, data *model.Attribute, id uint) (*model.Attribute, error) {
	if err := r.DB(ctx).
		Model(&model.Attribute{}).
		Where("attr_id = ?", id).
		Updates(data).Error; err != nil {
		return nil, err
	}

	return r.GetDetail(ctx, id)
}

// ================= 删除 =================

func (r *attributeRepository) Delete(ctx context.Context, id uint) error {
	result := r.DB(ctx).Where("attr_id = ?", id).Delete(&model.Attribute{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("id不存在")
	}

	return nil
}

// ================= 全部列表 =================

func (r *attributeRepository) GetList(ctx context.Context, id uint, q attributedto.RequestQuery) ([]model.Attribute, error) {
	var data []model.Attribute

	db := r.buildQuery(ctx, id, q)

	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// ================= 分页列表 =================

func (r *attributeRepository) GetPageList(ctx context.Context, id uint, q attributedto.RequestPageQuery) ([]model.Attribute, int64, error) {
	var data []model.Attribute

	db := r.buildQuery(ctx, id, q.RequestQuery)

	total, err := Paginate(db, &data, q.Page, q.PageSize)

	return data, total, err
}

// ================= 公共查询 =================
func (r *attributeRepository) buildQuery(ctx context.Context, id uint, q attributedto.RequestQuery) *gorm.DB {
	db := r.DB(ctx).Model(&model.Attribute{})

	if id > 0 {
		db = db.Where("cat_id = ?", id)
	}

	if q.Sel != "" {
		db = db.Where("attr_sel = ?", q.Sel)
	}

	return db
}

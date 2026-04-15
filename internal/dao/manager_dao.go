package dao

import (
	"context"
	"fmt"
	"go-server/internal/bootstrap"
	managerdto "go-server/internal/dto/manager"
	"go-server/internal/model"

	"gorm.io/gorm"
)

type ManagerRepository interface {
	GetByKeyWhere(ctx context.Context, name string) (*model.Manager, error)
	Create(ctx context.Context, user *model.Manager) (*model.Manager, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, info *model.Manager, id uint) (*model.Manager, error)
	GetDetail(ctx context.Context, id uint) (*model.Manager, error)
	GetLists(ctx context.Context, q managerdto.ManagerQuery) ([]model.Manager, int64, error)
}

func NewManagerRepository(
	r *bootstrap.Repository,
) ManagerRepository {
	return &managerRepository{
		Repository: r,
	}
}

type managerRepository struct {
	*bootstrap.Repository
}

// ================= 根据ID查询 =================

func (r *managerRepository) GetDetail(ctx context.Context, id uint) (*model.Manager, error) {
	var user model.Manager

	err := r.DB(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ================= 创建 =================

func (r *managerRepository) Create(ctx context.Context, user *model.Manager) (*model.Manager, error) {
	if err := r.DB(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// ================= 更新 =================

func (r *managerRepository) Update(ctx context.Context, user *model.Manager, id uint) (*model.Manager, error) {
	if err := r.DB(ctx).
		Model(&model.Manager{}).
		Where("mg_id = ?", id).
		Updates(user).Error; err != nil {
		return nil, err
	}

	return r.GetDetail(ctx, id)
}

// ================= 删除 =================

func (r *managerRepository) Delete(ctx context.Context, id uint) error {
	result := r.DB(ctx).Where("mg_id = ?", id).Delete(&model.Manager{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("用户不存在")
	}

	return nil
}

// ================= 根据管理员名查询 =================

func (r *managerRepository) GetByKeyWhere(ctx context.Context, name string) (*model.Manager, error) {
	var user model.Manager

	err := r.DB(ctx).Where("mg_name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ================= 管理员列表 分页 =================

func (r *managerRepository) GetLists(ctx context.Context, q managerdto.ManagerQuery) ([]model.Manager, int64, error) {
	var data []model.Manager

	db := r.buildQuery(ctx, q)

	// 分页
	total, err := Paginate(db, &data, q.Page, q.PageSize)

	return data, total, err
}

// ================= 公共查询 =================
func (r *managerRepository) buildQuery(ctx context.Context, q managerdto.ManagerQuery) *gorm.DB {
	db := r.DB(ctx).Model(&model.Manager{})

	if q.Query != nil {
		like := "%" + *q.Query + "%"
		db = db.Where("mg_name LIKE ?", like)
	}

	return db
}

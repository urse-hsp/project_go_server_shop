package dao

import (
	"context"
	"go-server/internal/bootstrap"
	"go-server/internal/model"
)

type ManagerRepository interface {
	GetManagerByUsername(ctx context.Context, username string) (*model.Manager, error)
	GetManagerLists(ctx context.Context, page, pageSize int) ([]model.Manager, int64, error)
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

// ================= 根据管理员名查询 =================

func (r *managerRepository) GetManagerByUsername(ctx context.Context, username string) (*model.Manager, error) {
	var user model.Manager

	err := r.DB(ctx).Where("mg_name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ================= 管理员列表 分页 =================

func (r *managerRepository) GetManagerLists(ctx context.Context, page, pageSize int) ([]model.Manager, int64, error) {
	var users []model.Manager

	db := r.DB(ctx).Model(&model.Manager{})

	// 分页
	total, err := Paginate(db, &users, page, pageSize)

	return users, total, err
}

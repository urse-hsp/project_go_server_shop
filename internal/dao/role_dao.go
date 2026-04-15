package dao

import (
	"context"
	"fmt"
	"go-server/internal/bootstrap"
	"go-server/internal/model"
)

type RoleRepository interface {
	GetByKeyWhere(ctx context.Context, username string) (*model.Role, error)

	Create(ctx context.Context, user *model.Role) (*model.Role, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, user *model.Role, id uint) (*model.Role, error)
	GetDetail(ctx context.Context, id uint) (*model.Role, error)
	GetList(ctx context.Context) ([]model.Role, error)
	GetLists(ctx context.Context, page, pageSize int) ([]model.Role, int64, error)
}

func NewRoleRepository(
	r *bootstrap.Repository,
) RoleRepository {
	return &roleRepository{
		Repository: r,
	}
}

type roleRepository struct {
	*bootstrap.Repository
}

// ================= 根据ID查询 =================

func (r *roleRepository) GetDetail(ctx context.Context, id uint) (*model.Role, error) {
	var user model.Role

	err := r.DB(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ================= 根据关键字查询 =================

func (r *roleRepository) GetByKeyWhere(ctx context.Context, username string) (*model.Role, error) {
	var user model.Role

	err := r.DB(ctx).Where("role_name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ================= 创建 =================

func (r *roleRepository) Create(ctx context.Context, user *model.Role) (*model.Role, error) {

	if err := r.DB(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// ================= 更新 =================

func (r *roleRepository) Update(ctx context.Context, user *model.Role, id uint) (*model.Role, error) {
	if err := r.DB(ctx).
		Model(&model.Role{}).
		Where("role_id = ?", id).
		Updates(user).Error; err != nil {
		return nil, err
	}

	return r.GetDetail(ctx, id)
}

// ================= 删除 =================

func (r *roleRepository) Delete(ctx context.Context, id uint) error {
	result := r.DB(ctx).Where("role_id = ?", id).Delete(&model.Role{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("用户不存在")
	}

	return nil
}

// ================= 全部列表 =================

func (r *roleRepository) GetList(ctx context.Context) ([]model.Role, error) {
	var list []model.Role

	if err := r.DB(ctx).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

// ================= 分页列表 =================

func (r *roleRepository) GetLists(ctx context.Context, page, pageSize int) ([]model.Role, int64, error) {
	var data []model.Role

	db := r.DB(ctx).Model(&model.Role{})

	// 分页
	total, err := Paginate(db, &data, page, pageSize)

	return data, total, err
}

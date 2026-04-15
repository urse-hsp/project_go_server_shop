package dao

import (
	"context"
	"fmt"
	"go-server/internal/bootstrap"
	"go-server/internal/dto/category"
	"go-server/internal/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(ctx context.Context, data *model.Category) (*model.Category, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, data *model.Category, id uint) (*model.Category, error)
	GetDetail(ctx context.Context, id uint) (*model.Category, error)
	GetList(ctx context.Context, q category.RequestQuery) ([]model.Category, error)
	GetPageList(ctx context.Context, q category.RequestPageQuery) ([]model.Category, int64, error)

	buildQuery(ctx context.Context, q category.RequestQuery) *gorm.DB
	GetByKeyWhere(ctx context.Context, username string) (*model.Category, error)
}

func NewCategoryRepository(
	r *bootstrap.Repository,
) CategoryRepository {
	return &categoryRepository{
		Repository: r,
	}
}

type categoryRepository struct {
	*bootstrap.Repository
}

// ================= 根据ID查询 =================

func (r *categoryRepository) GetDetail(ctx context.Context, id uint) (*model.Category, error) {
	var data model.Category

	err := r.DB(ctx).First(&data, id).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ================= 根据关键字查询 =================

func (r *categoryRepository) GetByKeyWhere(ctx context.Context, username string) (*model.Category, error) {
	var data model.Category

	err := r.DB(ctx).Where("cat_name = ?", username).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ================= 创建 =================

func (r *categoryRepository) Create(ctx context.Context, data *model.Category) (*model.Category, error) {

	if err := r.DB(ctx).Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// ================= 更新 =================

func (r *categoryRepository) Update(ctx context.Context, data *model.Category, id uint) (*model.Category, error) {
	if err := r.DB(ctx).
		Model(&model.Category{}).
		Where("cat_id = ?", id).
		Updates(data).Error; err != nil {
		return nil, err
	}

	return r.GetDetail(ctx, id)
}

// ================= 删除 =================

func (r *categoryRepository) Delete(ctx context.Context, id uint) error {
	result := r.DB(ctx).Where("cat_id = ?", id).Delete(&model.Category{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("id不存在")
	}

	return nil
}

// ================= 全部列表 =================

func (r *categoryRepository) GetList(ctx context.Context, q category.RequestQuery) ([]model.Category, error) {
	var data []model.Category

	db := r.buildQuery(ctx, q)

	if err := db.Where("cat_level = ? AND cat_deleted = 0", 0).Find(&data).Error; err != nil {
		return nil, err
	}

	if len(data) > 0 {
		if err := r.loadChildren(ctx, data, q.Type); err != nil {
			return nil, err
		}
	}

	return data, nil
}

// ================= 分页列表 =================

func (r *categoryRepository) GetPageList(ctx context.Context, q category.RequestPageQuery) ([]model.Category, int64, error) {
	var data []model.Category

	db := r.DB(ctx).Model(&model.Category{}).
		Where("cat_level = ? AND cat_deleted = 0", 0)

	total, err := Paginate(db, &data, q.Page, q.PageSize)
	if err != nil {
		return nil, 0, err
	}

	if len(data) > 0 {
		if err := r.loadChildren(ctx, data, q.Type); err != nil {
			return nil, 0, err
		}
	}

	return data, total, err
}

// ================= 公共查询 =================
func (r *categoryRepository) buildQuery(ctx context.Context, q category.RequestQuery) *gorm.DB {
	db := r.DB(ctx).Model(&model.Category{})

	if q.Query != nil {
		like := "%" + *q.Query + "%"
		db = db.Where("cat_name LIKE ?", like)
	}
	switch q.Type {
	case category.CategoryTypeLevel1:
		db = db.Where("cat_level = 0")

	case category.CategoryTypeLevel2:
		db = db.Where("cat_level <= 1")

	case category.CategoryTypeLevel3:
		db = db.Where("cat_level <= 2")
	}

	return db
}

// ================ 加载 children =================

func (r *categoryRepository) loadChildren(ctx context.Context, parents []model.Category, t category.CategoryType) error {

	// level1 不加载子级
	if t == category.CategoryTypeLevel1 {
		return nil
	}

	// 1️⃣ 收集一级ID
	var parentIDs []uint
	for _, v := range parents {
		parentIDs = append(parentIDs, v.CatID)
	}

	// 2️⃣ 查二级（Level2 和 Level3 都需要）
	var level2 []model.Category
	if err := r.DB(ctx).
		Where("cat_p_id IN ? AND cat_deleted = 0", parentIDs).
		Find(&level2).Error; err != nil {
		return err
	}

	// 👉 如果只是 Level2，到这里就可以返回
	if t == category.CategoryTypeLevel2 {
		level2Map := make(map[uint][]model.Category)
		for _, v := range level2 {
			level2Map[v.CatPID] = append(level2Map[v.CatPID], v)
		}

		for i := range parents {
			parents[i].Children = level2Map[parents[i].CatID]
		}

		return nil
	}

	// ================= Level3 =================

	// 3️⃣ 收集二级ID
	var level2IDs []uint
	for _, v := range level2 {
		level2IDs = append(level2IDs, v.CatID)
	}

	// 4️⃣ 查三级
	var level3 []model.Category
	if len(level2IDs) > 0 {
		if err := r.DB(ctx).
			Where("cat_p_id IN ? AND cat_deleted = 0", level2IDs).
			Find(&level3).Error; err != nil {
			return err
		}
	}

	// 5️⃣ 组装 三级 -> 二级
	level3Map := make(map[uint][]model.Category)
	for _, v := range level3 {
		level3Map[v.CatPID] = append(level3Map[v.CatPID], v)
	}

	for i := range level2 {
		level2[i].Children = level3Map[level2[i].CatID]
	}

	// 6️⃣ 组装 二级 -> 一级
	level2Map := make(map[uint][]model.Category)
	for _, v := range level2 {
		level2Map[v.CatPID] = append(level2Map[v.CatPID], v)
	}

	for i := range parents {
		parents[i].Children = level2Map[parents[i].CatID]
	}

	return nil
}

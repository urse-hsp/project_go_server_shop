package dao

import (
	"context"
	"fmt"
	"go-server/internal/bootstrap"
	goodsdto "go-server/internal/dto/goods"
	"go-server/internal/model"

	"gorm.io/gorm"
)

type GoodsRepository interface {
	GetByKeyWhere(ctx context.Context, username string) (*model.Goods, error)

	Create(ctx context.Context, user *model.Goods) (*model.Goods, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, user *model.Goods) (*model.Goods, error)
	GetDetail(ctx context.Context, id uint) (*model.Goods, error)
	GetList(ctx context.Context) ([]model.Goods, error)
	GetLists(ctx context.Context, q goodsdto.RequestPageQuery) ([]model.Goods, int64, error)
}

func NewGoodsRepository(
	r *bootstrap.Repository,
) GoodsRepository {
	return &goodsRepository{
		Repository: r,
	}
}

type goodsRepository struct {
	*bootstrap.Repository
}

// ================= 根据ID查询 =================

func (r *goodsRepository) GetDetail(ctx context.Context, id uint) (*model.Goods, error) {
	var data model.Goods

	err := r.DB(ctx).
		Preload("Attrs").
		Preload("Pics").First(&data, id).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ================= 根据关键字查询 =================

func (r *goodsRepository) GetByKeyWhere(ctx context.Context, username string) (*model.Goods, error) {
	var data model.Goods

	err := r.DB(ctx).Where("goods_name = ?", username).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// ================= 创建 =================

func (r *goodsRepository) Create(ctx context.Context, data *model.Goods) (*model.Goods, error) {
	// if err := r.DB(ctx).Create(data).Error; err != nil {
	// 	return nil, err
	// }

	// 事务
	err := r.DB(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(data).Error
	})
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ================= 更新 =================

// 先删再插
func (r *goodsRepository) Update(ctx context.Context, data *model.Goods) (*model.Goods, error) {
	err := r.DB(ctx).Transaction(func(tx *gorm.DB) error {

		// 1️⃣ 更新主表
		if err := tx.Model(&model.Goods{}).
			Where("goods_id = ?", data.GoodsID).
			Updates(data).Error; err != nil {
			return err
		}

		// 2️⃣ 删除旧的关联数据
		if err := tx.Where("goods_id = ?", data.GoodsID).
			Delete(&model.GoodsAttr{}).Error; err != nil {
			return err
		}

		if err := tx.Where("goods_id = ?", data.GoodsID).
			Delete(&model.GoodsPics{}).Error; err != nil {
			return err
		}

		// 3️⃣ 插入新的 Attrs
		if len(data.Attrs) > 0 {
			for i := range data.Attrs {
				data.Attrs[i].GoodsId = data.GoodsID
			}

			if err := tx.Create(&data.Attrs).Error; err != nil {
				return err
			}
		}

		// 4️⃣ 插入新的 Pics
		if len(data.Pics) > 0 {
			for i := range data.Pics {
				data.Pics[i].GoodsId = data.GoodsID
			}

			if err := tx.Create(&data.Pics).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return r.GetDetail(ctx, data.GoodsID)
}

// Association不可控容易出问题
func (r *goodsRepository) Update2(ctx context.Context, data *model.Goods) (*model.Goods, error) {
	err := r.DB(ctx).Transaction(func(tx *gorm.DB) error {

		// 1. 更新主表
		if err := tx.Model(&model.Goods{}).
			Where("goods_id = ?", data.GoodsID).
			Updates(data).Error; err != nil {
			return err
		}

		// 2. 设置外键
		for i := range data.Attrs {
			data.Attrs[i].GoodsId = data.GoodsID
		}
		for i := range data.Pics {
			data.Pics[i].GoodsId = data.GoodsID
		}

		// 3. 先清空
		if err := tx.Model(data).Association("Attrs").Clear(); err != nil {
			return err
		}

		if err := tx.Model(data).Association("Pics").Clear(); err != nil {
			return err
		}

		// 4. 再新增
		if len(data.Attrs) > 0 {
			if err := tx.Model(data).Association("Attrs").Append(data.Attrs); err != nil {
				return err
			}
		}

		if len(data.Pics) > 0 {
			if err := tx.Model(data).Association("Pics").Append(data.Pics); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return r.GetDetail(ctx, data.GoodsID)
}

// ================= 删除 =================

func (r *goodsRepository) Delete(ctx context.Context, id uint) error {
	result := r.DB(ctx).Where("goods_id = ?", id).Delete(&model.Goods{})

	// 逻辑删除
	// result := r.DB(ctx).Model(&model.Goods{}).Where("goods_id = ?", id).
	// 	Update("delete_time", time.Now().Unix())

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("不存在")
	}

	return nil
}

// ================= 全部列表 =================

func (r *goodsRepository) GetList(ctx context.Context) ([]model.Goods, error) {
	var data []model.Goods

	if err := r.DB(ctx).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// ================= 分页列表 =================

func (r *goodsRepository) GetLists(ctx context.Context, q goodsdto.RequestPageQuery) ([]model.Goods, int64, error) {
	var data []model.Goods

	db := r.buildQuery(ctx, q)

	// 分页
	total, err := Paginate(db, &data, q.Page, q.PageSize)

	return data, total, err
}

// ================= 公共查询 =================
func (r *goodsRepository) buildQuery(ctx context.Context, q goodsdto.RequestPageQuery) *gorm.DB {
	// 逻辑删除排除
	// db := r.DB(ctx).Model(&model.Goods{}).Where("delete_time = 0 OR delete_time IS NULL")

	db := r.DB(ctx).Model(&model.Goods{})

	if q.Query != nil {
		like := "%" + *q.Query + "%"
		db = db.Where("goods_name LIKE ?", like)
	}

	db = ApplySort(db, "goods_id", q.Sort)

	return db
}

// 通用dao
package dao

import (
	"fmt"
	goodsdto "go-server/internal/dto/goods"

	"gorm.io/gorm"
)

// 通用分页 page从1开始
func Paginate(db *gorm.DB, out interface{}, page, pageSize int) (int64, error) {
	var total int64

	// ⚠️ 注意：Count 要在 Offset/Limit 之前
	if err := db.Count(&total).Error; err != nil {
		return 0, err
	}

	err := db.
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(out).Error

	return total, err
}

func ApplySort(db *gorm.DB, field string, sort *goodsdto.Sort) *gorm.DB {
	order := "DESC"

	if sort != nil {
		switch *sort {
		case goodsdto.SortAsc:
			order = "ASC"
		case goodsdto.SortDesc:
			order = "DESC"
		}
	}

	return db.Order(fmt.Sprintf("%s %s", field, order))
}

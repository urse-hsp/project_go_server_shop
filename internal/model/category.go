package model

import "gorm.io/gorm"

// 商品分类表

type Category struct {
	CatID      uint           `gorm:"primaryKey;autoIncrement;comment:主键id"`
	CatName    string         `gorm:"type:varchar(255);default:null;comment:分类名称"`
	CatPID     uint           `gorm:"comment:分类父ID"`
	CatLevel   uint           `gorm:"comment:分类层级 0: 顶级 1:二级 2:三级"`
	CatDeleted uint           `gorm:"default:0;comment:是否删除 1为删除"`
	CatIcon    string         `gorm:"type:varchar(255);default:null;comment:分类图标"`
	CatSrc     string         `gorm:"type:text;default:null;comment:分类来源"`
	CgnCode    string         `gorm:"type:char(6);not null;default:'';comment:邮编"`
	DeleteTime gorm.DeletedAt `gorm:"default:null;comment:删除时间"`

	Children []Category `gorm:"-"`
}

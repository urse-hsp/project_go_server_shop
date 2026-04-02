package model

// 商品-分类关联表

type GoodsCats struct {
	CatsID     uint   `gorm:"primaryKey;autoIncrement;comment:主键"`
	ParentID   uint   `gorm:"not null;comment:父级id"`
	CatName    string `gorm:"type:varchar(50);not null;comment:分类名称"`
	IsShow     uint   `gorm:"not null;default:1;comment:是否显示"`
	CatSort    uint   `gorm:"not null;default:0;comment:分类排序"`
	DataFlag   uint   `gorm:"not null;default:1;comment:数据标记"`
	CreateTime int64  `gorm:"not null;comment:创建时间"`
}

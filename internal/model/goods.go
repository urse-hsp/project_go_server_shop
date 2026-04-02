package model

// 商品表

type Goods struct {
	GoodsID        uint    `gorm:"primaryKey;autoIncrement;comment:主键id"`
	GoodsName      string  `gorm:"type:varchar(255);not null;comment:商品名称"`
	GoodsPrice     float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:商品价格"`
	GoodsNumber    uint    `gorm:"not null;default:0;comment:商品数量"`
	GoodsWeight    uint    `gorm:"not null;default:0;comment:商品重量"`
	CatID          uint    `gorm:"not null;default:0;comment:类型id"`
	GoodsIntroduce string  `gorm:"type:text;comment:商品详情介绍"`
	GoodsBigLogo   string  `gorm:"type:char(128);not null;default:'';comment:图片logo大图"`
	GoodsSmallLogo string  `gorm:"type:char(128);not null;default:'';comment:图片logo小图"`
	IsDel          uint    `gorm:"not null;default:0;comment:0:正常  1:删除"`
	AddTime        int64   `gorm:"not null;comment:添加商品时间"`
	UpdTime        int64   `gorm:"not null;comment:修改商品时间"`
	DeleteTime     *int64  `gorm:"index;NULL;comment:软删除标志字段"`
	CatOneID       uint    `gorm:"default:0;comment:一级分类id"`
	CatTwoID       uint    `gorm:"default:0;comment:二级分类id"`
	CatThreeID     uint    `gorm:"default:0;comment:三级分类id"`
	HotMumber      uint    `gorm:"not null;default:0;comment:热卖数量"`
	IsPromote      uint    `gorm:"not null;default:0;comment:是否促销"`
	GoodsState     uint    `gorm:"not null;default:0;comment:商品状态 0: 未通过 1: 审核中 2: 已审核"`
}

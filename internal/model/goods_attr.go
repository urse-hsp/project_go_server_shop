package model

// 商品-属性关联表

type GoodsAttr struct {
	ID        uint    `gorm:"primaryKey;autoIncrement;comment:主键"`
	GoodsId   uint    `gorm:"not null;comment:商品id"`
	AttrId    uint    `gorm:"not null;comment:属性id"`
	AttrValue string  `gorm:"type:text;not null;comment:商品对应属性的值"`
	AddPrice  float64 `gorm:"type:decimal(8,2);comment:该属性需要额外增加的价钱"`
}

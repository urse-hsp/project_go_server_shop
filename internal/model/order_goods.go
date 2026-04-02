package model

// 商品订单关联表

type orderGoods struct {
	ID              uint    `gorm:"primaryKey;autoIncrement;comment:主键"`
	OrderId         uint    `gorm:"not null;comment:订单id"`
	GoodsId         uint    `gorm:"not null;comment:商品id"`
	GoodsPrice      float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:商品单价"`
	GoodsNumber     uint    `gorm:"not null;default:1;comment:购买单个商品数量"`
	GoodsTotalPrice float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:商品小计价格"`
}

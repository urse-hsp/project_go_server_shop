package goodsdto

type AttrDTO struct {
	ID        uint    `json:"id"`         // 主键
	GoodsId   uint    `json:"goods_id"`   // 商品id
	AttrId    uint    `json:"attr_id"`    // 属性id
	AttrValue string  `json:"attr_value"` // 商品对应属性的值
	AddPrice  float64 `json:"add_price"`  // 该属性需要额外增加的价钱
}

type CreateAttr struct {
	AttrId    uint   `json:"attr_id" binding:"required"`
	AttrValue string `json:"attr_value" binding:"required"`
}

package goodsdto

type AttrDTO struct {
	ID        uint    `json:"id"`
	GoodsId   uint    `json:"goods_id"`
	AttrId    uint    `json:"attr_id"`
	AttrValue string  `json:"attr_value"`
	AddPrice  float64 `json:"add_price"`
}

type CreateAttr struct {
	AttrId    uint   `json:"attr_id" binding:"required"`
	AttrValue string `json:"attr_value" binding:"required"`
}

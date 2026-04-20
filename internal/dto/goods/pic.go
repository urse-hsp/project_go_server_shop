package goodsdto

type PicDTO struct {
	PicsID  uint   `json:"pics_id"`      // 主键
	GoodsID uint   `json:"goods_id"`     // 商品id
	PicsBig string `json:"pics_big_url"` // 相册大图
	PicsMid string `json:"pics_mid_url"` // 相册中图
	PicsSma string `json:"pics_sma_url"` // 相册小图s
}

type CreatePics struct {
	Pic string `json:"pic" binding:"required"`
	Url string `json:"url" binding:"required"`
}

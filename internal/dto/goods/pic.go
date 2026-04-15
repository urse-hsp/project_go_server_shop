package goodsdto

type PicDTO struct {
	PicsID  uint   `json:"pics_id"`
	GoodsID uint   `json:"goods_id"`
	PicsBig string `json:"pics_big_url"`
	PicsMid string `json:"pics_mid_url"`
	PicsSma string `json:"pics_sma_url"`
}

type CreatePics struct {
	Pic string `json:"pic" binding:"required"`
	Url string `json:"url" binding:"required"`
}

package goodsdto

import (
	v1 "go-server/api/v1"
)

// ================= 请求 DTO =================

type CreateRequest struct {
	// gorm.Model
	GoodsName      string       `json:"goods_name" binding:"required"`
	GoodsPrice     float64      `json:"goods_price" binding:"required"`
	GoodsWeight    uint         `json:"goods_weight" binding:"required"`
	GoodsCat       string       `json:"goods_cat" binding:"required"`
	GoodsIntroduce string       `json:"goods_introduce" binding:"required"`
	GoodsNumber    uint         `json:"goods_number" binding:"required"`
	Attrs          []CreateAttr `json:"attrs" binding:"required"`
	Pics           []CreatePics `json:"pics" binding:"required"`
}

type UpdateRequest struct {
	CreateRequest
}

type Sort string

const (
	SortAsc  Sort = "asc"
	SortDesc Sort = "desc"
)

type RequestQuery struct {
	Query *string `form:"query"`
	Sort  *Sort   `form:"sort"`
}
type RequestPageQuery struct {
	v1.PageRequest
	RequestQuery
}

// ================= 响应 DTO =================

// 对外公开（别人能看到）
type PublicDTO struct {
	GoodsID     uint    `json:"goods_id"`
	GoodsName   string  `json:"goods_name"`
	GoodsNumber uint    `json:"goods_number"`
	GoodsPrice  float64 `json:"goods_price"`
	GoodsState  uint    `json:"goods_state"`
	GoodsWeight uint    `json:"goods_weight"`
	IsPromote   uint    `json:"is_promote"`
	AddTime     uint    `json:"add_time"`
	// HotMumber uint    `json:"hot_mumber"`
}

type DetailPublicDTO struct {
	PublicDTO
	Attrs []AttrDTO `json:"attrs"`
	Pics  []PicDTO  `json:"pics"`

	GoodsIntroduce string `json:"goods_introduce"`
	CatId          uint   `json:"cat_id"`
	CatOneID       uint   `json:"cat_one_id"`
	CatTwoID       uint   `json:"cat_two_id"`
	CatThreeID     uint   `json:"cat_three_id"`
	GoodsBigLogo   string `json:"goods_big_logo"`
	GoodsCat       string `json:"goods_cat"`
	GoodsSmallLogo string `json:"goods_small_logo"`

	UpdTime uint `json:"upd_time"`
}

// 私有（自己能看到）
type PrivateDTO struct {
}

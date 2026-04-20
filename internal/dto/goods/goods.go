package goodsdto

import (
	v1 "go-server/api/v1"
)

// ================= 请求 DTO =================

type CreateRequest struct {
	// gorm.Model
	GoodsName      string       `json:"goods_name" binding:"required"`      // 商品名称
	GoodsPrice     float64      `json:"goods_price" binding:"required"`     // 商品价格
	GoodsWeight    uint         `json:"goods_weight" binding:"required"`    // 商品重量
	GoodsCat       string       `json:"goods_cat" binding:"required"`       // 商品分类
	GoodsIntroduce string       `json:"goods_introduce" binding:"required"` // 商品介绍
	GoodsNumber    uint         `json:"goods_number" binding:"required"`    // 商品数量
	Attrs          []CreateAttr `json:"attrs" binding:"required"`           // 商品属性
	Pics           []CreatePics `json:"pics" binding:"required"`            // 商品图片
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
	Query *string `form:"query"` // 搜索
	Sort  *Sort   `form:"sort"`  // 排序
}
type RequestPageQuery struct {
	v1.PageRequest
	RequestQuery
}

// ================= 响应 DTO =================
type PageResponse struct {
	Data []PublicDTO `json:"data"` // 列表
	v1.PageSizeResponse
}

// 对外公开（别人能看到）
type PublicDTO struct {
	GoodsID     uint    `json:"goods_id"`     // 主键
	GoodsName   string  `json:"goods_name"`   // 商品名称
	GoodsNumber uint    `json:"goods_number"` // 商品数量
	GoodsPrice  float64 `json:"goods_price"`  // 商品价格
	GoodsState  uint    `json:"goods_state"`  // 商品状态
	GoodsWeight uint    `json:"goods_weight"` // 商品重量
	IsPromote   uint    `json:"is_promote"`   // 是否促销
	AddTime     uint    `json:"add_time"`     // 添加时间s
	// HotMumber uint    `json:"hot_mumber"`
}

type DetailPublicDTO struct {
	PublicDTO
	Attrs []AttrDTO `json:"attrs"` // 商品属性
	Pics  []PicDTO  `json:"pics"`  // 商品图片

	GoodsIntroduce string `json:"goods_introduce"`  // 商品介绍
	CatId          uint   `json:"cat_id"`           // 商品分类
	CatOneID       uint   `json:"cat_one_id"`       // 商品分类
	CatTwoID       uint   `json:"cat_two_id"`       // 商品分类
	CatThreeID     uint   `json:"cat_three_id"`     // 商品分类
	GoodsBigLogo   string `json:"goods_big_logo"`   // 商品大图
	GoodsCat       string `json:"goods_cat"`        // 商品分类
	GoodsSmallLogo string `json:"goods_small_logo"` // 商品小图

	UpdTime uint `json:"upd_time"` // 更新时间
}

// 私有（自己能看到）
type PrivateDTO struct {
}

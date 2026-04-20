package attributedto

import (
	v1 "go-server/api/v1"
	"go-server/internal/model"
)

// ================= 请求 DTO =================
type CreateRequest struct {
	AttrName string        `json:"attr_name" binding:"required"` // 属性名称
	Sel      model.AttrSel `json:"attr_sel" binding:"required"`  // 属性录入方式
}

type UpdateRequest struct {
	CreateRequest
	AttrVals *string `json:"attr_vals" binding:"omitempty"`
}

type RequestQuery struct {
	Sel model.AttrSel `form:"sel" binding:"omitempty"` // 属性录入方式
}

type RequestPageQuery struct {
	v1.PageRequest
	RequestQuery
}

// ================= 响应 DTO =================

// 对外公开（别人能看到）
type PublicDTO struct {
	AttrID    uint            `json:"attr_id"`    // 属性id
	AttrName  string          `json:"attr_name"`  // 属性名称
	CatID     uint            `json:"cat_id"`     // 分类id
	AttrSel   model.AttrSel   `json:"attr_sel"`   // 属性录入方式
	AttrWrite model.AttrWrite `json:"attr_write"` // 属性录入方式
	AttrVals  string          `json:"attr_vals"`  // 可选值
}

// 私有（自己能看到）
type PrivateDTO struct {
}

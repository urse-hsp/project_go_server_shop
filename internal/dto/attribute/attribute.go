package attributedto

import (
	v1 "go-server/api/v1"
	"go-server/internal/model"
)

// ================= 请求 DTO =================
type CreateRequest struct {
	// gorm.Model
	AttrName string        `json:"attr_name"`
	Sel      model.AttrSel `json:"attr_sel"`
}

type UpdateRequest struct {
	CreateRequest
	AttrVals *string `json:"attr_vals"`
}

type RequestQuery struct {
	// Query *string `form:"query"`
	Sel model.AttrSel `form:"sel"`
}
type RequestPageQuery struct {
	v1.PageRequest
	RequestQuery
}

// ================= 响应 DTO =================

// 对外公开（别人能看到）
type PublicDTO struct {
	AttrID    uint            `json:"attr_id"`
	AttrName  string          `json:"attr_name"`
	CatID     uint            `json:"cat_id"`
	AttrSel   model.AttrSel   `json:"attr_sel"`
	AttrWrite model.AttrWrite `json:"attr_write"`
	AttrVals  string          `json:"attr_vals"`
}

// 私有（自己能看到）
type PrivateDTO struct {
}

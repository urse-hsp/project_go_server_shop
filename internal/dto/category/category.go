package category

import v1 "go-server/api/v1"

type CategoryType string

const (
	CategoryTypeLevel1 CategoryType = "1"
	CategoryTypeLevel2 CategoryType = "2"
	CategoryTypeLevel3 CategoryType = "3"
)

// ================= 请求 DTO =================
type CreateRequest struct {
	CatName  string `json:"cat_name" binding:"required"` // 分类名称
	CatPID   uint   `json:"cat_pid"`                     // 父ID
	CatLevel uint   `json:"cat_level"`                   // 分类层级
}

type UpdateRequest struct {
	CatName string `json:"cat_name" binding:"omitempty"` // 分类名称
}

type PageRequest struct {
	Page     *int `form:"current"`
	PageSize *int `form:"pageSize"`
}

type RequestQuery struct {
	Query *string      `form:"query"`
	Type  CategoryType `form:"type"`
	// Type *string `form:"type" binding:"omitempty,oneof=1 2 3"`
	PageRequest
}

type RequestPageQuery struct {
	v1.PageRequest
	RequestQuery
}

// ================= 响应 DTO =================

// 对外公开（别人能看到）
type PublicDTO struct {
	CatsID     uint        `json:"cat_id"`             // 分类ID
	CatLevel   uint        `json:"cat_level"`          // 分类层级
	CatName    string      `json:"cat_name"`           // 分类名称
	CatPID     uint        `json:"cat_pid"`            // 父ID
	CatDeleted bool        `json:"cat_deleted"`        // 是否删除
	Children   []PublicDTO `json:"children,omitempty"` // omitempty：如果字段是“空值”，就不返回
}

// 私有（自己能看到）
type PrivateDTO struct {
}

package category

import (
	v1 "go-server/api/v1"
)

type CategoryType string

const (
	CategoryTypeLevel1 CategoryType = "1"
	CategoryTypeLevel2 CategoryType = "2"
	CategoryTypeLevel3 CategoryType = "3"
)

// ================= 请求 DTO =================
type CreateRequest struct {
	// gorm.Model
	CatName  string `json:"cat_name" binding:"required"`
	CatPID   uint   `json:"cat_pid"`
	CatLevel uint   `json:"cat_level"`
}

type UpdateRequest struct {
	CatName string `json:"cat_name" binding:"omitempty"`
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
	CatsID     uint        `json:"cat_id"`
	CatLevel   uint        `json:"cat_level"`
	CatName    string      `json:"cat_name"`
	CatPID     uint        `json:"cat_pid"`
	CatDeleted bool        `json:"cat_deleted"`
	Children   []PublicDTO `json:"children,omitempty"` // omitempty：如果字段是“空值”，就不返回
}

// 私有（自己能看到）
type PrivateDTO struct {
}

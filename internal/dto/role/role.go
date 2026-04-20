package roledto

// ================= 请求 DTO =================
type LoginRequest struct {
	RoleName string `json:"roleName" binding:"required"`
	RoleDesc string `json:"roleDesc" binding:"required"`
}

// ================= 响应 DTO =================

type RolePublicDTO struct {
	ID       uint   `json:"id"`       // 角色ID
	RoleName string `json:"roleName"` // 角色名称
	RoleDesc string `json:"roleDesc"` // 角色描述
	children []any  // 子级
}

package roledto

// ================= 请求 DTO =================
type LoginRequest struct {
	RoleName string `json:"roleName" binding:"required"`
	RoleDesc string `json:"roleDesc" binding:"required"`
}

// ================= 响应 DTO =================

type RolePublicDTO struct {
	ID       uint   `json:"id"`
	RoleName string `json:"roleName"`
	RoleDesc string `json:"roleDesc"`
	children []any
}

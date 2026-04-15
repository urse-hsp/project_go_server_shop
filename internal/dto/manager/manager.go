package managerdto

import v1 "go-server/api/v1"

// ================= 请求 DTO =================

// 登录
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 注册
type CreateRequest struct {
	LoginRequest
	Email  string `json:"email" binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
}

type UpdateRequest struct {
	Email  *string `json:"email"`
	Mobile *string `json:"mobile"`
	State  *bool   `json:"state"`
}

type ManagerQuery struct {
	Query *string `form:"query"`
	v1.PageRequest
}

// ================= 响应 DTO =================

// 👉 对外公开（别人能看到）
type ManagerPublicDTO struct {
	ID   uint   `json:"id"`
	User string `json:"username"`
}

// 私有（自己能看到）
type ManagerPrivateDTO struct {
	MgID      uint   `json:"id"`
	MgName    string `json:"username"`
	MgEmail   string `json:"email"`
	MgMobile  string `json:"mobile"`
	MgState   bool   `json:"mg_state"`
	RoleID    uint   `json:"role_id"`
	Role_name string `json:"role_name"`
}

// ================= 登录返回 =================

type LoginResponse struct {
	ManagerPrivateDTO
	Token string `json:"token"`
}

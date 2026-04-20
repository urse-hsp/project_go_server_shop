package managerdto

import v1 "go-server/api/v1"

// ================= 请求 DTO =================

// 登录
type LoginRequest struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

// 注册
type CreateRequest struct {
	LoginRequest
	Email  string `json:"email" binding:"required"`  // 邮箱
	Mobile string `json:"mobile" binding:"required"` // 手机号
}

type UpdateRequest struct {
	Email  *string `json:"email"`  // 邮箱
	Mobile *string `json:"mobile"` // 手机号
	State  *bool   `json:"state"`  // 状态
}

type ManagerQuery struct {
	Query *string `form:"query"` // 搜索
	v1.PageRequest
}

// ================= 响应 DTO =================

type PageResponse struct {
	Data []ManagerPublicDTO `json:"data"` // 列表
	v1.PageSizeResponse
}

// 对外公开（别人能看到）
type ManagerPublicDTO struct {
	ID   uint   `json:"id"`       // ID
	User string `json:"username"` // 用户名
}

// 私有（自己能看到）
type ManagerPrivateDTO struct {
	MgID      uint   `json:"id"`        // ID
	MgName    string `json:"username"`  // 用户名
	MgEmail   string `json:"email"`     // 邮箱
	MgMobile  string `json:"mobile"`    // 手机号
	MgState   bool   `json:"mg_state"`  // 状态
	RoleID    uint   `json:"role_id"`   // 角色id
	Role_name string `json:"role_name"` // 角色
}

// ================= 登录返回 =================

type LoginResponse struct {
	ManagerPrivateDTO
	Token string `json:"token"`
}

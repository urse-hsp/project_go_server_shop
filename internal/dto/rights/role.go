package rightsdto

// ================= 请求 DTO =================

// ================= 响应 DTO =================

type PublicDTO struct {
	ID      uint   `json:"id"`       // 权限id
	PsName  string `json:"authName"` // 权限名称
	PsC     string `json:"path"`     // 权限路径
	PsPid   uint   `json:"pid"`      // 父id
	PsLevel string `json:"level"`    // 权限等级
}

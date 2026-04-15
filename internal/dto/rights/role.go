package rightsdto

// ================= 请求 DTO =================

// ================= 响应 DTO =================

type PublicDTO struct {
	ID      uint   `json:"id"`
	PsName  string `json:"authName"`
	PsC     string `json:"path"`
	PsPid   uint   `json:"pid"`
	PsLevel string `json:"level"`
}

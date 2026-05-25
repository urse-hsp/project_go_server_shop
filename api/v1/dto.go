package v1

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

type PageRequest struct {
	Page     int `form:"page" binding:"omitempty,min=1"`          // 可选，最小为1
	PageSize int `form:"limit" binding:"omitempty,min=1,max=999"` // 可选，1~100
}

type PageInfo struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

type PageResponse[T any] struct {
	Data []T `json:"data"` // 列表
	PageInfo
}

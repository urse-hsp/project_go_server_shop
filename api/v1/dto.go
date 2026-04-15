package v1

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

type PageRequest struct {
	Page     int `form:"current" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=1,max=100"`
}

// type PageRequest2 struct {
// 	*PageRequest
// }

type PageResponse struct {
	Data     any `json:"data"`
	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

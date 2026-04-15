package orderdto

import (
	v1 "go-server/api/v1"
	"go-server/internal/model"
)

// ================= 请求 DTO =================
type CreateRequest struct {
	// gorm.Model
	Username string
}

type IsSend string

const (
	IsSendYes IsSend = "1"
	IsSendNo  IsSend = "0"
)

type UpdateRequest struct {
	OrderPrice *float64         `json:"order_price"`
	OrderPay   *model.OrderPay  `json:"order_pay"`
	IsSend     *IsSend          `json:"is_send"`
	PayStatus  *model.PayStatus `json:"pay_status"`
}

type RequestQuery struct {
	Query *string `form:"query"`
}
type RequestPageQuery struct {
	v1.PageRequest
	RequestQuery
}

// ================= 响应 DTO =================

// 对外公开（别人能看到）
type PublicDTO struct {
	OrderID            uint            `json:"order_id"`
	UserID             uint            `json:"user_id"`
	OrderNumber        string          `json:"order_number"`
	OrderPrice         float64         `json:"order_price"`
	OrderPay           model.OrderPay  `json:"order_pay"`
	IsSend             model.IsSend    `json:"is_send"`
	TradeNo            string          `json:"trade_no"`
	OrderFapiaoTitle   string          `json:"order_fapiao_title"`
	OrderFapiaoCompany string          `json:"order_fapiao_company"`
	OrderFapiaoContent string          `json:"order_fapiao_content"`
	ConsigneeAddr      string          `json:"consignee_addr"`
	PayStatus          model.PayStatus `json:"pay_status"`
	CreateTime         int64           `json:"create_time"`
	UpdateTime         int64           `json:"update_time"`
}

// 私有（自己能看到）
type PrivateDTO struct {
}

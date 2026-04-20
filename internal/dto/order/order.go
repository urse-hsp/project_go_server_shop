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
	OrderPrice *float64         `json:"order_price"` // 订单价格
	OrderPay   *model.OrderPay  `json:"order_pay"`   // 支付方式
	IsSend     *IsSend          `json:"is_send"`     // 是否发货
	PayStatus  *model.PayStatus `json:"pay_status"`  // 支付状态
}

type RequestQuery struct {
	Query *string `form:"query"`
}
type RequestPageQuery struct {
	v1.PageRequest
	RequestQuery
}

// ================= 响应 DTO =================

type PageResponse struct {
	Data []PublicDTO `json:"data"` // 列表
	v1.PageSizeResponse
}

// 对外公开（别人能看到）
type PublicDTO struct {
	OrderID            uint            `json:"order_id"`             // 订单id
	UserID             uint            `json:"user_id"`              // 用户id
	OrderNumber        string          `json:"order_number"`         // 订单号
	OrderPrice         float64         `json:"order_price"`          // 订单价格
	OrderPay           model.OrderPay  `json:"order_pay"`            // 支付方式
	IsSend             model.IsSend    `json:"is_send"`              // 是否发货
	TradeNo            string          `json:"trade_no"`             // 交易号
	OrderFapiaoTitle   string          `json:"order_fapiao_title"`   // 发票抬头
	OrderFapiaoCompany string          `json:"order_fapiao_company"` // 发票公司
	OrderFapiaoContent string          `json:"order_fapiao_content"` // 发票内容
	ConsigneeAddr      string          `json:"consignee_addr"`       // 收货地址
	PayStatus          model.PayStatus `json:"pay_status"`           // 支付状态
	CreateTime         int64           `json:"create_time"`          // 创建时间
	UpdateTime         int64           `json:"update_time"`          // 更新时间
}

// 私有（自己能看到）
type PrivateDTO struct {
}

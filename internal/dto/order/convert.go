package orderdto

import "go-server/internal/model"

// ================= DTO 转换 =================

// 👉 他人可见
func ToPublicDTO(u *model.Order) PublicDTO {
	return PublicDTO{
		OrderID:            u.OrderID,
		UserID:             u.UserID,
		OrderNumber:        u.OrderNumber,
		OrderPrice:         u.OrderPrice,
		OrderPay:           u.OrderPay,
		IsSend:             u.IsSend,
		TradeNo:            u.TradeNo,
		OrderFapiaoTitle:   u.OrderFapiaoTitle,
		OrderFapiaoCompany: u.OrderFapiaoCompany,
		OrderFapiaoContent: u.OrderFapiaoContent,
		ConsigneeAddr:      u.ConsigneeAddr,
		PayStatus:          u.PayStatus,
		CreateTime:         u.CreateTime,
		UpdateTime:         u.UpdateTime,
	}
}

// 👉 自己可见
func ToPrivateDTO(u *model.Order) PrivateDTO {
	return PrivateDTO{}
}

func ListToPublic(users []model.Order) []PublicDTO {
	list := make([]PublicDTO, 0, len(users))
	for _, u := range users {
		list = append(list, ToPublicDTO(&u))
	}
	return list
}

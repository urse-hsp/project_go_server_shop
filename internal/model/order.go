package model

// 订单表

type order struct {
	OrderID            uint    `gorm:"primaryKey;autoIncrement;comment:主键"`
	UserID             uint    `gorm:"not null;comment:下订单会员id"`
	OrderNumber        string  `gorm:"type:varchar(32);not null;comment:订单编号"`
	OrderPrice         float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:订单总金额"`
	OrderPay           string  `gorm:"type:enum('0','1','2','3');not null;default:'1';comment:支付方式  0未支付 1支付宝  2微信  3银行卡"`
	IsSend             string  `gorm:"type:enum('是','否');not null;default:'否';comment:订单是否已经发货"`
	TradeNo            string  `gorm:"type:varchar(32);not null;default:'';comment:支付宝交易流水号码"`
	OrderFapiaoTitle   string  `gorm:"type:enum('个人','公司');not null;default:'个人';comment:发票抬头 个人 公司"`
	OrderFapiaoCompany string  `gorm:"type:varchar(32);not null;default:'';comment:公司名称"`
	OrderFapiaoContent string  `gorm:"type:varchar(32);not null;default:'';comment:发票内容"`
	ConsigneeAddr      string  `gorm:"type:text;not null;comment:consignee收货人地址"`
	PayStatus          string  `gorm:"type:enum('0','1');not null;default:'0';comment:订单状态： 0未付款、1已付款"`
	CreateTime         int64   `gorm:"not null;comment:记录生成时间"`
	UpdateTime         int64   `gorm:"not null;comment:记录修改时间"`
}

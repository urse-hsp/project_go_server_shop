package model

// 快递表

type Express struct {
	ExpressID  uint   `gorm:"primaryKey;autoIncrement;comment:主键id"`
	OrderID    uint   `gorm:"not null;comment:订单id"`
	ExpressCom string `gorm:"type:varchar(32);comment:订单快递公司名称"`
	ExpressNu  string `gorm:"type:varchar(32);comment:快递单编号"`
	CreateTime int64  `gorm:"not null;comment:记录生成时间"`
	UpdateTime int64  `gorm:"not null;comment:记录修改时间"`
}

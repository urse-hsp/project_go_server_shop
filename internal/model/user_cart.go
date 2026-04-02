package model

// 用户购物车模型定义

type UserCart struct {
	CartID uint `gorm:"primaryKey;autoIncrement;comment:主键"`
	UserID uint `gorm:"not null;comment:用户ID"`
	// CardInfo   string `gorm:"type:text;comment:购物车信息，JSON格式"`
	CartInfo   string `gorm:"type:json;comment:购物车信息(JSON)"`
	CreateTime int64  `gorm:"autoCreateTime;NULL;comment:注册时间"`
	UpdateTime int64  `gorm:"autoUpdateTime;NULL;comment:更新时间"`
	DeleteTime int64  `gorm:"index;NULL;comment:删除时间"`
}

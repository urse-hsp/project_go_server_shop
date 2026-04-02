package model

// 收货人表

type Consignee struct {
	CgnID      uint   `gorm:"primaryKey;autoIncrement;comment:主键id"`
	UserID     uint   `gorm:"not null;comment:会员id"`
	CgnName    string `gorm:"type:varchar(32);not null;comment:收货人名称"`
	CgnAddress string `gorm:"type:varchar(200);not null;default:'';comment:收货人地址"`
	CgnTel     string `gorm:"type:varchar(20);not null;default:'';comment:收货人电话"`
	CgnCode    string `gorm:"type:char(6);not null;default:'';comment:邮编"`
	DeleteTime *int64 `gorm:"default:null;comment:删除时间"`
}

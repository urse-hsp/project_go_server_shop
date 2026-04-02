package model

// 类型表

type Type struct {
	TypeID     uint   `gorm:"primaryKey;autoIncrement;comment:主键"`
	TypeName   string `gorm:"type:varchar(32);not null;comment:类型名称"`
	DeleteTime int64  `gorm:"index;NULL;comment:删除时间"`
}

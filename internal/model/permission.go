package model

// 权限表

type Permission struct {
	ID         uint   `gorm:"primaryKey;autoIncrement;comment:主键"`
	PsName     string `gorm:"type:varchar(20);not null;comment:权限名称"`
	PsPid      uint   `gorm:"not null;comment:父id"`
	PsC        string `gorm:"type:varchar(32);not null;default:'';comment:控制器"`
	PsA        string `gorm:"type:varchar(32);not null;default:'';comment:操作方法"`
	PsLevel    string `gorm:"type:enum('0','2','1');not null;default:'0';comment:权限等级"`
	DeleteTime int64  `gorm:"index;NULL;comment:删除时间"`
}

package model

//

type Role struct {
	RoleID     uint   `gorm:"primaryKey;autoIncrement;comment:主键"`
	RoleName   string `gorm:"type:varchar(32);not null;comment:类型名称"`
	Ids        string `gorm:"type:varchar(512);not null;default:'';comment:权限ids,1,2,5"`
	Ca         string `gorm:"type:text;comment:控制器-操作,控制器-操作,控制器-操作"`
	RoleDesc   string `gorm:"type:text;comment:角色描述"`
	DeleteTime int64  `gorm:"index;NULL;comment:删除时间"`
}

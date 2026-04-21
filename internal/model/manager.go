package model

// 管理员表

type Manager struct {
	MgID     uint   `gorm:"primaryKey;autoIncrement;comment:主键"`
	MgName   string `gorm:"type:varchar(32);not null;comment:管理员名称"`
	MgPwd    string `gorm:"type:char(64);not null;comment:密码"`
	MgTime   int    `gorm:"not null;comment:注册时间"`
	RoleID   uint   `gorm:"not null;default:0;comment:角色id"`
	MgMobile string `gorm:"type:varchar(32);comment:手机号"`
	MgEmail  string `gorm:"type:varchar(64);comment:邮箱"`
	MgState  uint   `gorm:"not null;default:1;comment:1：表示启用 0:表示禁用"`
}

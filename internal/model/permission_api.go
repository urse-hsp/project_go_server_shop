package model

type permissionApi struct {
	ID           uint   `gorm:"primaryKey;autoIncrement;comment:主键"`
	PsID         uint   `gorm:"not null;comment:权限ID"`
	PsApiService string `gorm:"type:varchar(255);comment:API服务"`
	PsApiAction  string `gorm:"type:varchar(255);comment:API操作"`
	PsApiPath    string `gorm:"type:varchar(255);comment:API路径"`
	PsApiOrder   int    `gorm:"comment:API排序"`
}

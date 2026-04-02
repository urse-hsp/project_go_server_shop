package model

type report_1 struct {
	ID           uint   `gorm:"primaryKey;autoIncrement;comment:主键"`
	Rp1UserCount int    `gorm:"comment:用户数"`
	Rp1Area      string `gorm:"type:varchar(128);comment:地区"`
	Rp1Date      string `gorm:"type:date;comment:日期"`
}

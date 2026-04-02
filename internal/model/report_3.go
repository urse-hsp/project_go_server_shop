package model

type report_3 struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;comment:主键"`
	Rp3Src   string `gorm:"type:varchar(127);comment:用户来源"`
	Rp3Count int    `gorm:"comment:数量"`
	Rp3Date  string `gorm:"type:datetime;comment:日期"`
}

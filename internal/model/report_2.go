package model

type report_2 struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;comment:主键"`
	Rp2Page  string `gorm:"type:varchar(128);comment:页面"`
	Rp2Count int    `gorm:"comment:访问量"`
	Rp2Date  string `gorm:"type:date;comment:日期"`
}

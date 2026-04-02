package model

// 属性表

type Attribute struct {
	AttrID     uint   `gorm:"primaryKey;autoIncrement;comment:主键id"`
	AttrName   string `gorm:"type:varchar(32);not null;comment:属性名称"`
	CatID      uint   `gorm:"not null;comment:外键，类型id"`
	AttrSel    string `gorm:"type:enum('only','many');not null;default:'only';comment:only:输入框(唯一)  many:后台下拉列表/前台单选框"`
	AttrWrite  string `gorm:"type:enum('manual','list');not null;default:'manual';comment:manual:手工录入  list:从列表选择"`
	AttrVals   string `gorm:"type:text;null;comment:可选值列表信息,例如颜色：白色,红色,绿色,多个可选值通过逗号分隔"`
	DeleteTime *int64 `gorm:"default:null;comment:删除时间标志"`
}

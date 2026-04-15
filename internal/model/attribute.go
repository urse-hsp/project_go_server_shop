package model

// ================= 枚举类型 =================

// 属性选择类型
type AttrSel string

const (
	AttrSelOnly AttrSel = "only"
	AttrSelMany AttrSel = "many"
)

// 属性录入方式
type AttrWrite string

const (
	AttrWriteManual AttrWrite = "manual"
	AttrWriteList   AttrWrite = "list"
)

// ================= 模型 =================

type Attribute struct {
	AttrID     uint      `gorm:"primaryKey;autoIncrement;comment:主键id"`
	AttrName   string    `gorm:"type:varchar(32);not null;comment:属性名称"`
	CatID      uint      `gorm:"not null;comment:外键，分类id"`
	AttrSel    AttrSel   `gorm:"type:enum('only','many');not null;default:'only';comment:only:输入框(唯一) many:多选"`
	AttrWrite  AttrWrite `gorm:"type:enum('manual','list');not null;default:'manual';comment:manual:手工录入 list:从列表选择"`
	AttrVals   string    `gorm:"type:text;comment:可选值(逗号分隔)"`
	DeleteTime *int64    `gorm:"default:null;comment:软删除时间"`
}

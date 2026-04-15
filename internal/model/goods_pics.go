package model

// 商品-相册关联表

type GoodsPics struct {
	PicsID  uint   `gorm:"primaryKey;autoIncrement;comment:主键"`
	GoodsId uint   `gorm:"not null;comment:商品id"`
	PicsBig string `gorm:"type:char(128);not null;default:'';comment:相册大图800*800"`
	PicsMid string `gorm:"type:char(128);not null;default:'';comment:相册中图350*350"`
	PicsSma string `gorm:"type:char(128);not null;default:'';comment:相册小图50*50"`
}

package model

// 用户模型定义

type IsActive string

const (
	IsActiveYes IsActive = "是"
	IsActiveNo  IsActive = "否"
)

type UserSex string

const (
	UserSexMale   UserSex = "男"
	UserSexFemale UserSex = "女"
	UserSexSecret UserSex = "保密"
)

type UserXueli string

const (
	UserXueliBoshi    UserXueli = "博士"
	UserXueliShuoshi  UserXueli = "硕士"
	UserXueliBenke    UserXueli = "本科"
	UserXueliZhuanke  UserXueli = "专科"
	UserXueliGaozhong UserXueli = "高中"
	UserXueliChuzhong UserXueli = "初中"
	UserXueliXiaoxue  UserXueli = "小学"
)

type User struct {
	ID            uint      `gorm:"primaryKey;autoIncrement;comment:自增ID"`
	Username      string    `gorm:"type:varchar(128);not null;unique;comment:用户名"`
	QQOpenID      string    `gorm:"type:char(32);not null;unique;comment:qq官方唯一编号信息"`
	Password      string    `gorm:"type:char(64);not null;comment:密码"`
	UserEmail     string    `gorm:"type:varchar(64);not null;unique;comment:电子邮件"`
	UserEmailCode string    `gorm:"type:char(13); null;comment:新用户注册邮件激活唯一校验码'"`
	IsActive      IsActive  `gorm:"type:enum('是','否');not null;default:'否';comment:新用户是否已经通过邮箱激活帐号"`
	UserSex       UserSex   `gorm:"type:enum('保密','女','男');not null;default:'男';comment:性别"`
	UserQQ        string    `gorm:"type:varchar(32);not null;unique;comment:qq号码"`
	UserTel       string    `gorm:"type:varchar(32);not null;unique;comment:手机号"`
	UserXueli     UserXueli `gorm:"type:enum('博士','硕士','本科','专科','高中','初中','小学');default:'本科';not null;comment:学历"`
	UserHobby     string    `gorm:"type:varchar(255);not null;comment:爱好"`
	UserIntroduce string    `gorm:"text;comment:个人简介"`
	CreateTime    int64     `gorm:"autoCreateTime;not null;comment:注册时间"`
	UpdateTime    int64     `gorm:"autoUpdateTime;not null;comment:更新时间"`
}

// // gorm 默认表名是结构体名的复数形式，重写 TableName 方法可以自定义表名
// func (User) TableName() string {
// 	return "users"
// }

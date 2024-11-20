package model

// Role 角色
type Role struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(50);not null;default:'';comment:角色名称" json:"name"`
	Code string `gorm:"column:code;type:varchar(50);unique;not null;default:'';comment:角色代码" json:"code"`
	Menu Menu   // 直接嵌入 Menu 结构体
}

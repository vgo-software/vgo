package role

import (
	"ych/vgo/app/common"
)

// RoleMenu 角色菜单
type RoleMenu struct {
	common.BaseModel
	RoleId uint64 `gorm:"column:role_id;not null;default:0;comment:角色ID" json:"role_id"`
	MenuId uint64 `gorm:"column:menu_id;not null;default:0;comment:菜单ID" json:"menu_id"`
}

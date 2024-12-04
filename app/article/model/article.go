package model

import "vgo-software/vgo/app/common/model"

type Article struct {
	model.BaseModel
	Title   string `gorm:"column:title;type:varchar(255);not null;default:'';comment:标题;" validate:"required,min=2" json:"title"`
	Status  int    `gorm:"column:status;type:smallint;not null;default:1;comment:状态 (1启用 2禁用)" json:"status"`
	Content string `gorm:"column:content;type:text;comment:公告内容;" json:"content"`
}

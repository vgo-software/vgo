package model

import "vgo-software/vgo/app/common/model"

type Module struct {
	model.BaseModel
	Name string `gorm:"column:name;type:varchar(255);not null;default:'';comment:name;" validate:"required" json:"name"`
}

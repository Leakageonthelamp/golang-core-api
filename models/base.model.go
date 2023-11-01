package models

import (
	"time"

	"github.com/Leakageonthelamp/go-leakage-core/utils"
)

type BaseModel struct {
	ID        string     `json:"id" gorm:"column:id; primary_key"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at; type:timestamp without time zone; not null; default:now()"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at; type:timestamp without time zone; not null; default:now()"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"column:delete_at; type:timestamp without time zone; default:null"`
}

func NewBaseModel() BaseModel {
	return BaseModel{
		ID:        utils.GetUUID(),
		CreatedAt: utils.GetCurrentDateTime(),
		UpdatedAt: utils.GetCurrentDateTime(),
	}
}

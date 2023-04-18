package models

import (
	
	"gorm.io/gorm"
)

type Role_Permissions struct {
	gorm.Model
	UserId int `json:"user_id" gorm:"size:30;not null`
	MenuId string `json:"menu_id" gorm:"size:30;not null"`
	ModuleId string `json:"module_id" gorm:"size:30;"`
}




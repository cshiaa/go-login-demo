package models

import (
	
	"gorm.io/gorm"

	"github.com/cshiaa/go-login-demo/global"

)

type RolePermissions struct {
	gorm.Model
	UserId uint `json:"user_id" gorm:"size:30;not null`
	MenuId string `json:"menu_id" gorm:"size:30;not null"`
	ModuleId string `json:"module_id" gorm:"size:30;"`
}


func (role *RolePermissions) Insert(uid uint, menu Menu) error {

	role.UserId = uid
	if menu.ParentId == "0" {
		role.MenuId = menu.MenuId
		role.ModuleId = menu.ParentId
	} else {
		role.MenuId = menu.ParentId
		role.ModuleId = menu.MenuId
	}

	if err := global.RY_DB.Debug().Create(&role).Error; err!= nil{
		return err
	}
	
	return nil
}

func (role *RolePermissions) Delete(uid uint, menu Menu) error {

	role.UserId = uid
	if menu.ParentId == "0" {
		if err := global.RY_DB.Debug().Where("menu_id = ?", menu.MenuId).Delete(&role).Error; err!= nil{
			return err
		}
	} else {
		if err := global.RY_DB.Debug().Where("module_id = ?", menu.MenuId).Delete(&role).Error; err!= nil{
			return err
		}
	}


	
	return nil
}

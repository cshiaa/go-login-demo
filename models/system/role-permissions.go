package system

import (
	
	"gorm.io/gorm"

	"github.com/cshiaa/go-login-demo/global"

)

type RolePermissions struct {
	gorm.Model
	UserId uint `json:"user_id" gorm:"size:30;not null`
	MenuId string `json:"menu_id" gorm:"size:30;not null"`
	// ModuleId string `json:"module_id" gorm:"size:30;"`
}


func (role *RolePermissions) Insert(uid uint, menu Menu) error {

	role.UserId = uid
	role.MenuId = menu.MenuId

	if err := global.RY_DB.Create(&role).Error; err!= nil{
		return err
	}

	return nil
}

func (role *RolePermissions) Delete(uid uint, menu Menu) error {

	role.UserId = uid
	if err := global.RY_DB.Where("menu_id = ?", menu.MenuId).Delete(&role).Error; err!= nil{
		return err
	}

	return nil
}



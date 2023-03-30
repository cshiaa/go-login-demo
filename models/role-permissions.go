package models

import (
	
	"gorm.io/gorm"
	"github.com/cshiaa/go-login-demo/global"
)

type Role_Permissions struct {
	gorm.Model
	UserId int `json:"user_id" gorm:"size:30;not null`
	MenuId string `json:"menu_id" gorm:"size:30;not null"`
	ModuleId string `json:"module_id" gorm:"size:30;"`
}



//用户是否含有子菜单
func IsChildMenu(uid uint) (bool, error) {

	if err := global.RY_DB.Where("user_id =? AND module_id is not null", uid).Find(&Role_Permissions{}).Error; err!= nil {
        return false, err
    }
	return true, nil
}



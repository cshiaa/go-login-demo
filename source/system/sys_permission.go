package system

import (

	"github.com/cshiaa/go-login-demo/global"
	systemModel "github.com/cshiaa/go-login-demo/models/system"
)

func GetUserChildRoleMenuID(uid uint) ([]string, error) {

	var childMenuIds []string
	subQuery := global.RY_DB.Model(&systemModel.Menu{}).Select("menu_id").Where("parent_id != 0")
	if err := global.RY_DB.Model(&systemModel.RolePermissions{}).Where("user_id = ? and menu_id in (?)", uid, subQuery).Pluck("menu_id", &childMenuIds).Error; err != nil {
        return nil,err
    }
	
	return childMenuIds,nil
}

func GetUserRoleMenuID(uid uint) ([]string, error) {

	var menuIds []string
	if err := global.RY_DB.Model(&systemModel.RolePermissions{}).Where("user_id = ?", uid).Pluck("menu_id", &menuIds).Error; err != nil {
        return nil,err
    }

	return menuIds,nil
}
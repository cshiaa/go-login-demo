package system

import (
	"errors"

	"github.com/cshiaa/go-login-demo/global"
	systemModel "github.com/cshiaa/go-login-demo/models/system"
)

type PermissionMenu struct {
	Menuname 		string 				`json:"menuName"`
	MenuId 			string 				`json:"menuId"`
	ParentId 		string     			`json:"parentId"`
	Children 		[]PermissionMenu 	`json:"children" gorm:"-"`
}

func GetMenuObject(m []string)(menu []systemModel.Menu, err error) {

	if err := global.RY_DB.Where("menu_id in (?)", m).Find(&menu).Error; err!= nil{
		return menu, errors.New("该用户没有找到对应的菜单")
	}

	return menu, nil
}

//获取已有的全部菜单
func GetAllMenu(parentId string) (menus []PermissionMenu, err error) {

	if err := global.RY_DB.Model(&systemModel.Menu{}).Where("parent_id = ?", parentId).Select("menu_id", "menuname", "parent_id").Find(&menus).Error; err != nil {
        return nil,err
    }

	for index, menu := range menus {
		menus[index].Children, _ = GetAllMenu(menu.MenuId)
	}
	return menus, nil
}

//获取用户菜单
func GetUserMenu(uid uint, parentId string) (menus []systemModel.Menu, err error){

	subQuery := global.RY_DB.Model(&systemModel.RolePermissions{}).Select("menu_id").Where("user_id = ?", uid)
	if err := global.RY_DB.Where("parent_id = ? and menu_id in (?)", parentId, subQuery).Find(&menus).Error; err!= nil{
		return menus, err
	}

	for index, menu := range menus {
		menus[index].Children, _ = GetUserMenu(uid, menu.MenuId)
	}
	return menus, nil

}
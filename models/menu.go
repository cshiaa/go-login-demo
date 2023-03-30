package models

import (
	"errors"
	"gorm.io/gorm"

	"github.com/cshiaa/go-login-demo/global"
)

type Menu struct {
	gorm.Model
	Icon string 	`json:"icon" gorm:"size:30;not null;unique"`
	Menuname string `json:"menuname" gorm:"size:255;not null;`
	MenuUrl string 		`json:"url" gorm:"size:30;not null; default:''"`
	MenuType int 		`json:"menutype" gorm:"size:4; not null;default:0"`
	MenuId string 		`json:"menuId" gorm:"size:10; not null`
	ParentId string     `json:"parentId" gorm:"size:10; "`
}

type Menus struct {
	Menu
	Children []Menu	`json:"children"`
}

func GetMenuID(uid uint) (menu []Menu, err error) {

	// var menu Menu

	if err := global.RY_DB.Where("menu_type = ?", "0").Find(&menu).Error; err != nil {
		return menu, err
	}
	//menus := DB.Exec("SElECT id, icon, menuname, url FROM menus")


	return menu, nil

}

//获取用户菜单
func getUserPartentMenu(uid uint) (menu []Menu, err error) {

	subQuery := global.RY_DB.Select("menu_id").Where("user_id = ?  and module_id is null", uid).Table("role_permissions")
	if err := global.RY_DB.Where("menu_id in (?)", subQuery).Find(&menu).Error; err!= nil{
		return menu, errors.New("User not found Permissions")
	}

	return menu, nil

}

//获取用户子菜单
func getUserChildMenu(uid uint, menuId string) (menu []Menu, err error) {


	subQuery := global.RY_DB.Select("module_id").Where("user_id = ? and menu_id = ? and module_id is not null", uid, menuId).Table("role_permissions")
	if err := global.RY_DB.Where("menu_id in (?)", subQuery).Find(&menu).Error; err!= nil{
		return menu, errors.New("User not found Child Menu")
	}
	return menu, nil
}

func GetMenu(uid uint) (menus []Menus, err error){

	userMenus := make([]Menus, 0)

	partentMenu, err := getUserPartentMenu(uid)
	if err != nil {
		return userMenus, errors.New("User not found Partent Menu")
	}
	for _, menu := range partentMenu{
		
		childSlice := make([]Menu, 0)
		childMenu, err := getUserChildMenu(uid, menu.MenuId)
		if err != nil {
			return userMenus, errors.New("User not found Children Menu")
		}
		if childMenu != nil {
			childSlice = append(childSlice, childMenu...)
		}
		menuTmp := Menus{
			Menu: menu,
			Children: childSlice,
		}
		userMenus = append(userMenus, menuTmp)
		// fmt.Println("menu: ", menu)
		// fmt.Println("childSlice: ", childSlice)
		// fmt.Println(userMenus)
		// menuTmp.FormatJSON()
		// jsonbyte,_ := json.Marshal(userMenus)
		// fmt.Println(string(jsonbyte))
	}
	return userMenus, nil
}
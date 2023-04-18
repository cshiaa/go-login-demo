package models

import (
	"errors"
	"gorm.io/gorm"

	"github.com/cshiaa/go-login-demo/global"
)


type Menu struct {
	gorm.Model
	Icon string 	`json:"icon" gorm:"size:30;unique"`
	Menuname string `json:"menuname" gorm:"size:255;not null;`
	MenuUrl string 		`json:"url" gorm:"size:30;not null; default:''"`
	MenuType int 		`json:"menutype" gorm:"size:4; not null;default:0"`
	MenuId string 		`json:"menuId" gorm:"size:10; not null`
	ParentId string     `json:"parentId" gorm:"size:10; not null"`
	Name string 	`json:"name" gorm:"size:20; not null"`
}

type Menus struct {
	Menu
	Children []Menu	`json:"children"`
}

type PermissionMenu struct {
	Menuname string `json:"menuName"`
	MenuId string 	`json:"menuId"`
}

type PermissionMenus struct {
	PermissionMenu
	Children []PermissionMenu `json:"children"`
}

//判断是否为二级菜单
func (menu *Menu)IsChildrenMenu() bool {
	return menu.MenuType != 0
}

//判断是否为一级菜单
func (menu *Menu)IsParentMenu() bool {
	return menu.MenuType != 1
}


func GetMenuObject(m []string)(menu []Menu, err error) {

	if err := global.RY_DB.Where("menu_id in (?)", m).Find(&menu).Error; err!= nil{
		return menu, errors.New("该用户没有找到对应的菜单")
	}

	return menu, nil
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

func GetUserMenu(uid uint) (menus []Menus, err error){

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


//获取一级菜单
func getPartentMenu() (menu []PermissionMenu, err error) {

	if err := global.RY_DB.Table("menus").Where("menu_type = 0 ").Find(&menu).Error; err!= nil{
		return menu, errors.New("not found Permissions")
	}

	return menu, nil

}

//获取一级菜单的子菜单
func getChildMenu(menuId string) (menu []PermissionMenu, err error) {


	if err := global.RY_DB.Table("menus").Where("menu_type = 1 and parent_id = (?)", menuId).Find(&menu).Error; err!= nil{
		return menu, errors.New("not found Child Menu")
	}
	return menu, nil
}


func GetMenu() (menus []PermissionMenus, err error){

	userMenus := make([]PermissionMenus, 0)

	partentMenu, err := getPartentMenu()
	if err != nil {
		return userMenus, errors.New("not found Partent Menu")
	}
	for _, menu := range partentMenu{

		childSlice := make([]PermissionMenu, 0)
		childMenu, err := getChildMenu(menu.MenuId)
		if err != nil {
			return userMenus, errors.New("not found Children Menu")
		}
		if childMenu != nil {
			childSlice = append(childSlice, childMenu...)
		}
		menuTmp := PermissionMenus{
			PermissionMenu: menu,
			Children: childSlice,
		}
		userMenus = append(userMenus, menuTmp)
	}
	return userMenus, nil
}


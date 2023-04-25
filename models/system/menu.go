package system

import (
	"gorm.io/gorm"

)


type Menu struct {
	gorm.Model
	Icon 			string 				`json:"icon" gorm:"size:30;unique"`
	Menuname 		string 				`json:"menuname" gorm:"size:255;not null;`
	MenuUrl 		string 				`json:"url" gorm:"size:30;not null; default:''"`
	MenuType 		int 				`json:"menutype" gorm:"size:4; not null;default:0"`
	MenuId 			string 				`json:"menuId" gorm:"size:10; not null`
	ParentId 		string     			`json:"parentId" gorm:"size:10; not null"`
	Name 			string 				`json:"name" gorm:"size:20; not null"`
	// Meta          						`json:"meta" gorm:"embedded;comment:附加属性"`                            // 附加属性
	Children      	[]Menu          	`json:"children" gorm:"-"`
}

type Meta struct {
	Title       string `json:"title" gorm:"comment:菜单名"`                       // 菜单名
}

//判断是否为二级菜单
func (menu *Menu)IsChildrenMenu() bool {
	return menu.MenuType != 0
}

//判断是否为一级菜单
func (menu *Menu)IsParentMenu() bool {
	return menu.MenuType != 1
}





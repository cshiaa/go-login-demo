package controller

import (
	"net/http"
	"strconv"

	"github.com/cshiaa/go-login-demo/models"
	"github.com/gin-gonic/gin"

	"github.com/cshiaa/go-login-demo/utils"
	"github.com/cshiaa/go-login-demo/utils/tools"
)

type NewMenuPermissions struct {
	Menus []string	`json:"menusList"`
}

//根据请求中携带的token获取用户菜单列表
func GetMenuList(c *gin.Context){

	userId, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menu, err := models.GetUserMenu(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","menuList": menu})
}

//获取目前所有的菜单列表
func GetAllMenu(c *gin.Context) {

	menu, err := models.GetMenu()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","menuList": menu})

}

//对用户拥有的菜单列表格式化处理，只返回已拥有的菜单权限的menuid列表
//这里只返回了子菜单的权限列表
func getMenuId(menu []models.Menus) (idList []string, err error) {
	
	for _, m := range menu {
		// idList = append(idList, m.MenuId)
		if len(m.Children) > 0 {
			for _, childm := range m.Children {
				idList = append(idList, childm.MenuId)
			}
		}
	}
	return idList, nil
}

//获取用户的菜单ID
func getUserMenuId(uid uint) (menuId []string, err error) {


	menu, err := models.GetUserMenu(uid)
	if err != nil {
		return nil, err
	}

	menuId, err = getMenuId(menu)
	if err != nil {
		return nil, err
	}
	return menuId, nil
}

//获取用户的菜单列表
func GetUserMenuList(c *gin.Context) {

	uid := c.Query("userid")
	uidInt, _ := strconv.Atoi(uid)
	uidUint := uint(uidInt)
	
	userMenuIdList, err :=  getUserMenuId(uidUint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"success","userMenuList": userMenuIdList})

}

//更新用户的菜单权限
func UpdateUserMenu(c *gin.Context) {

	uid := c.Query("userid")
	uidInt, _ := strconv.Atoi(uid)
	uidUint := uint(uidInt)

	// var newUserMenus  = []string{"5", "4", "1", "2"}
	var newMenu = NewMenuPermissions{}
	if err := c.ShouldBindJSON(&newMenu); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	// newMenus, err := models.GetMenuObject(menus)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }


	srcUserMenus, err :=  getUserMenuId(uidUint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addMenus, delMenus := tools.Arrcmp[string](srcUserMenus, newMenu.Menus)
	
	addMenuObjectList, _ := models.GetMenuObject(addMenus)
	for _, m := range addMenuObjectList {
		var role = models.RolePermissions{}
		err := role.Insert(uidUint, m)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	delMenuObjectList, _ := models.GetMenuObject(delMenus)
	for _, m := range delMenuObjectList {
		var role = models.RolePermissions{}
		err := role.Delete(uidUint, m)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","addMenuList": addMenus, "delMenuList": delMenus, "newUserMenus":newMenu })

}
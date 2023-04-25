package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/cshiaa/go-login-demo/utils"
	"github.com/cshiaa/go-login-demo/utils/tools"
	systemModel "github.com/cshiaa/go-login-demo/models/system"
	"github.com/cshiaa/go-login-demo/source/system"
)

type NewMenuPermissions struct {
	Menus []string	`json:"menusList"`
}

//一级菜单父ID为0
var partenId string = "0"

//根据请求中携带的token获取用户菜单列表
func GetMenuList(c *gin.Context){

	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menu, err := system.GetUserMenu(uid, partenId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","menuList": menu})
}

//获取目前所有的菜单列表
func GetMenu(c *gin.Context) {

	menu, err := system.GetAllMenu(partenId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","menuList": menu})

}

//获取用户的菜单列表
func GetUserMenuList(c *gin.Context) {

	uid := c.Query("userid")
	uidInt, _ := strconv.Atoi(uid)
	uidUint := uint(uidInt)
	
	userChildMenus, err :=  system.GetUserChildRoleMenuID(uidUint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"success","userChildMenus": userChildMenus})

}

//更新用户的菜单权限
func UpdateUserMenu(c *gin.Context) {

	uid := c.Query("userid")
	uidInt, _ := strconv.Atoi(uid)
	uidUint := uint(uidInt)

	var newMenu = NewMenuPermissions{}
	if err := c.ShouldBindJSON(&newMenu); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	userMenuIds, err := system.GetUserRoleMenuID(uidUint)
	srcUserMenus := tools.DuplicateRemovingMap[string](userMenuIds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addMenus, delMenus := tools.Arrcmp[string](srcUserMenus, newMenu.Menus)

	addMenuObjectList, _ := system.GetMenuObject(addMenus)
	for _, m := range addMenuObjectList {
		var role = systemModel.RolePermissions{}
		err := role.Insert(uidUint, m)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	delMenuObjectList, _ := system.GetMenuObject(delMenus)
	for _, m := range delMenuObjectList {
		var role = systemModel.RolePermissions{}
		err := role.Delete(uidUint, m)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message":"success" })

}
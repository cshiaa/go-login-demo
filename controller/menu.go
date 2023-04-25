package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/cshiaa/go-login-demo/utils"
	"github.com/cshiaa/go-login-demo/utils/tools"
	systemModel "github.com/cshiaa/go-login-demo/models/system"
	"github.com/cshiaa/go-login-demo/source/system"
	"github.com/cshiaa/go-login-demo/common/response"
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
		response.FailWithMessage("Token解析用户失败", c)
		return
	}

	menu, err := system.GetUserMenu(uid, partenId)
	if err != nil {
		response.FailWithMessage("用户获取菜单列表失败", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","menuList": menu})
}

//获取目前所有的菜单列表
func GetMenu(c *gin.Context) {

	menu, err := system.GetAllMenu(partenId)
	if err != nil {
		response.FailWithMessage("获取所有菜单列表失败", c)
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
		response.FailWithMessage("用户获取子菜单失败", c)
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
		response.FailWithMessage("解析请求更新的菜单列表报错", c)
        return
	}

	userMenuIds, err := system.GetUserRoleMenuID(uidUint)
	if err != nil {
		response.FailWithMessage("获取用户菜单ID列表失败", c)
		return
	}

	srcUserMenus := tools.DuplicateRemovingMap[string](userMenuIds)
	addMenus, delMenus := tools.Arrcmp[string](srcUserMenus, newMenu.Menus)

	addMenuObjectList, _ := system.GetMenuObject(addMenus)
	for _, m := range addMenuObjectList {
		var role = systemModel.RolePermissions{}
		err := role.Insert(uidUint, m)
		if err != nil {
			response.FailWithMessage("用户添加菜单权限失败", c)
			return
		}
	}

	delMenuObjectList, _ := system.GetMenuObject(delMenus)
	for _, m := range delMenuObjectList {
		var role = systemModel.RolePermissions{}
		err := role.Delete(uidUint, m)
		if err != nil {
			response.FailWithMessage("用户删除菜单权限失败", c)
			return
		}
	}

	response.SuccessWithMessage("更新菜单权限成功", c)

}
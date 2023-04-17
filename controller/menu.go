package controller

import (
	"net/http"
	"strconv"

	"github.com/cshiaa/go-login-demo/models"
	"github.com/gin-gonic/gin"

	"github.com/cshiaa/go-login-demo/utils"
)


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


func GetAllMenu(c *gin.Context) {

	menu, err := models.GetMenu()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","menuList": menu})

}

func getMenuId(menu []models.Menus) (idList []string, err error) {
	
	for _, m := range menu {
		idList = append(idList, m.MenuId)
		if len(m.Children) > 0 {
			for _, childm := range m.Children {
				idList = append(idList, childm.MenuId)
			}
		}
	}
	return idList, nil
}

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

//检查提交的用户权限，返回没有的权限
// func checkUserPermissions(uid uint, authMenu []string)(addMenu []string, delMenu []string, err error){
// 	currentUserMenu, err := getUserMenuId(uid)

// }



func UpdateUserMenu(c *gin.Context) {

	uid := c.Query("userid")
	uidInt, _ := strconv.Atoi(uid)
	uidUint := uint(uidInt)

	menu, err := models.GetUserMenu(uidUint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userMenuIdList, err := getMenuId(menu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"success","userMenuList": userMenuIdList})

}
package system

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/cshiaa/go-login-demo/utils"
	"github.com/cshiaa/go-login-demo/models/system"
	"github.com/cshiaa/go-login-demo/global"

)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (atoken, rtoken string, err error) {

	u := system.User{}
	err = global.RY_DB.Model(system.User{}).Where("username = ?", username).Take(&u).Error
	if err!= nil {
		return "", "", err
	}

	err = VerifyPassword(password, u.Password)
	if err!= nil {
        return "", "", err
    }

	atoken, rtoken, err = utils.GenToken(u.ID, u.Username)
	if err != nil {
		return "", "", err
	}

	return atoken, rtoken, nil
}

func GetUserByID(uid uint) (u system.User, err error) {

	if err := global.RY_DB.First(&u,uid).Error; err != nil {
		return u,errors.New("User not found!")
	}

	u.PrepareGive()
	
	return u,nil

}

func GetAllUser() (users []system.UserInfo, err error) {
	if err := global.RY_DB.Table("users").Select("id", "username").Order("id").Find(&users).Error; err != nil {
		return users,errors.New("User not found!")
	}
	return users,nil
}
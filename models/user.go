package models


import (
	"html"
	"strings"
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"github.com/cshiaa/go-login-demo/utils"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:30;not null;unique"`
	Password string `json:"password" gorm:"size:255;not null;`
}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err!= nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) BeforeSave() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err!= nil {
		return err
	}

	u.Password = string(hashedPassword)

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (atoken, rtoken string, err error) {

	// var err error

	u := User{}

	err = DB.Model(User{}).Where("username =?", username).Take(&u).Error
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

func GetUserByID(uid uint) (User,error) {

	var u User

	if err := DB.First(&u,uid).Error; err != nil {
		return u,errors.New("User not found!")
	}

	u.PrepareGive()
	
	return u,nil

}

func (u *User) PrepareGive(){
	u.Password = ""
}
package system


import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/cshiaa/go-login-demo/global"
)

type User struct {
	gorm.Model
	Username string `json:"username, omitempty" gorm:"size:30;not null;unique"`
	Password string `json:"password, omitempty" gorm:"size:255;not null;`
}

type UserInfo struct {
	ID        	uint `json:"id"`
	Username 	string `json:"username, omitempty"`
}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = global.RY_DB.Create(&u).Error
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

func (u *User) PrepareGive(){
	u.Password = ""
}
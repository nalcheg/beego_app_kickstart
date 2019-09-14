package models

import (
	"github.com/anacrolix/log"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "users"
}

type User struct {
	Id       int
	Login    string
	Password string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func CheckLoginPassword(login, password string) bool {
	var user User
	o := orm.NewOrm()

	err := o.QueryTable("users").Filter("Login", login).One(&user)
	if err != nil {
		log.Printf("%v", err)

		return false
	}

	if user.Id > 0 {
		if CheckPasswordHash(password, user.Password) == true {

			return true
		}
	}

	return false
}

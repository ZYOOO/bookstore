package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

var user = &model.User{
	Username: "JoyBoy12",
	Password: "123321",
	Email:    "ZYOOOOO@126.com",
}

//func TestUser(t *testing.T) {
//	//fmt.Println("测试userDao中的函数")
//	//t.Run("测试注册", testRegister)
//	//t.Run("测试登陆", testLogin)
//}

func TestLogin(t *testing.T) {
	user1, _ := CheckUsernameAndPassword(user.Username+"asdasd", user.Password)
	if user1 != nil {
		fmt.Println(user1)
	} else {
		fmt.Println("uname or psw error")
	}
}

func TestRegister(t *testing.T) {
	if !UsernameExisted(user.Username) {
		SaveUser(user)
	} else {
		fmt.Println("username already existed")
	}
}

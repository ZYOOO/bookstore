package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

//登陆, 注册, 保存(注册成功后的) 三个方法

// CheckUsernameAndPassword 用户登录之后返回一个User
func CheckUsernameAndPassword(username, password string) (*model.User, error) {
	sql := "select id,username,password,email from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sql, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

// UsernameExisted 注册之前先查看此用户名是否存在
func UsernameExisted(username string) bool {
	sql := "select id from users where username = ?"
	row := utils.Db.QueryRow(sql, username)
	id := -1
	row.Scan(&id)
	if id != -1 {
		return true
	}
	return false
}

// SaveUser 可以注册, 存储用户
func SaveUser(user *model.User) (bool, error) {
	sql := "insert into users(username,password,email) values(?,?,?)"
	_, err := utils.Db.Exec(sql, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return false, err
	}
	return true, nil
}

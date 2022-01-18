package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"html/template"
	"net/http"
)

// Login 用户登录处理器
func Login(w http.ResponseWriter, r *http.Request) {
	//获取请求中的用户登录信息
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//验证
	user, _ := dao.CheckUsernameAndPassword(username, password)
	if user.ID > 0 {
		//正确就跳转
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确!")
	}
}

// Register 用户注册处理器
func Register(w http.ResponseWriter, r *http.Request) {
	//获取请求中的用户登录信息
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//验证
	if dao.UsernameExisted(username) {
		//正确就跳转
		t := template.Must(template.ParseFiles("views/pages/user/register.html"))
		t.Execute(w, "用户名已存在!")
	} else {
		user := &model.User{
			Username: username,
			Password: password,
			Email:    email,
		}
		dao.SaveUser(user)
		t := template.Must(template.ParseFiles("views/pages/user/register_success.html"))
		t.Execute(w, "")
	}
}

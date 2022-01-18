package controller

import (
	"bookstore/dao"
	"fmt"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//获取请求中的用户登录信息
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//验证
	user, _ := dao.CheckUsernameAndPassword(username, password)
	fmt.Println("获取的用户是:", user)
	if user.ID > 0 {
		//正确就跳转
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

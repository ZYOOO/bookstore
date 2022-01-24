package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"html/template"
	"net/http"
)

// Login 用户登录处理器
func Login(w http.ResponseWriter, r *http.Request) {
	//判断用户是否已经登陆
	if flag, _ := IsLogin(r); flag {
		GetPageBooksByPrice(w, r)
		return
	}
	//获取请求中的用户登录信息
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//验证
	user, _ := dao.CheckUsernameAndPassword(username, password)
	if user.ID > 0 {
		//正确就跳转
		//生成uuid作为sessionID
		uuid := utils.CreateUUID()
		sess := &model.Session{
			SessionID: uuid,
			Username:  user.Username,
			UserID:    user.ID,
		}
		dao.AddSession(sess)
		cookie := http.Cookie{
			Name:     "user",
			Value:    uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, user)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确!")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		dao.DeleteSession(cookieValue)
		//设置cookie失效 0:未设置 <0:立即失效 >0:秒数
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	GetPageBooksByPrice(w, r)
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

// CheckUsername 通过发送Ajax请求验证用户名是否存在
func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	if dao.UsernameExisted(username) {
		w.Write([]byte("用户名已存在!"))
	} else {
		w.Write([]byte("<font style='color:green'>用户名可用!</font>"))
	}
}

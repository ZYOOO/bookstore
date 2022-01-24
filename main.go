package main

import (
	"bookstore/controller"
	"net/http"
)

func main() {
	//处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	//引导页
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	//去登陆页面
	http.HandleFunc("/login", controller.Login)
	//注销
	http.HandleFunc("/logout", controller.Logout)
	//去注册页面
	http.HandleFunc("/register", controller.Register)
	//Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUsername", controller.CheckUsername)
	//获取带页面的图书
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	//根据BID删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//修改图书
	http.HandleFunc("/editBook", controller.EditBook)
	//更新或添加的图书
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)
	//价格查询
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	//开启对端口8080的监听
	http.ListenAndServe(":8080", nil)
}

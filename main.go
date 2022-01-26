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
	//添加图书到购物车
	http.HandleFunc("/addBookToCart", controller.AddBookToCart)
	//获取购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	//清空购物车
	http.HandleFunc("/clearCart", controller.ClearCart)
	//删除单项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	//更新购物项
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	//去结账
	http.HandleFunc("/checkout", controller.Checkout)
	//获取所有订单
	http.HandleFunc("/getOrders", controller.GetOrders)
	//订单详情
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	//获取我的订单
	http.HandleFunc("/getMyOrder", controller.GetMyOrders)
	//发货
	http.HandleFunc("/sendOrder", controller.SendOrder)
	//确认收货
	http.HandleFunc("/takeOrder", controller.TakeOrder)
	//开启对端口8080的监听
	http.ListenAndServe(":8080", nil)
}

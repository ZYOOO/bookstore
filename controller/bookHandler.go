package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

// IndexHandler 跳转首页
//func IndexHandler(w http.ResponseWriter, r *http.Request) {
//	//获取页码
//	pageNo := r.FormValue("pageNo")
//	if pageNo == "" {
//		pageNo = "1"
//	}
//	page, _ := dao.GetPageBooks(pageNo)
//	t := template.Must(template.ParseFiles("views/index.html"))
//	t.Execute(w, page)
//}

// GetPageBooks  获取带分页的图书,后台查看用的,前台用ByPrice
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	//这里是不是可以不用模板,直接用w.Write然后前端对data进行处理, 嫁鸡随鸡嫁狗随狗吧
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//同时跳转页面
	t.Execute(w, page)
}

func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	min := r.FormValue("min")
	max := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}
	page := &model.Page{}
	if min == "" && max == "" {
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		page, _ = dao.GetPageBooksByPrice(pageNo, min, max)
		page.MinPrice = min
		page.MaxPrice = max
	}
	//检测用户是否登陆
	if ok, sess := IsLogin(r); ok {
		//查到了, 已经登陆
		page.IsLogin = true
		page.Username = sess.Username
	}
	//这里是不是可以不用模板,直接用w.Write然后前端对data进行处理, 嫁鸡随鸡嫁狗随狗吧
	t := template.Must(template.ParseFiles("views/index.html"))
	//同时跳转页面
	t.Execute(w, page)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bID := r.FormValue("bID")
	dao.DeleteBook(bID)
	GetPageBooks(w, r)
}

// EditBook 去更新或者添加图书的页面
func EditBook(w http.ResponseWriter, r *http.Request) {
	bID := r.FormValue("bID")
	book, _ := dao.GetBookByBID(bID)
	if book.BID > 0 {
		//更新图书
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
		//在添加图书
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}
}

// UpdateOrAddBook 更新或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	bID := r.PostFormValue("bID")
	//数据格式转换
	iBID, _ := strconv.ParseInt(bID, 10, 0)
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	book := &model.Book{
		BID:     iBID,
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   iSales,
		Stock:   iStock,
		ImgPath: "/static/img/default.jpg",
	}
	if book.BID > 0 {
		//更新
		dao.UpdateBook(book)
		//调用处理器重新查询一次数据库之后跳转
	} else {
		//添加
		dao.AddBook(book)
	}
	GetPageBooks(w, r)
}

func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		sess, _ := dao.GetSession(cookieValue)
		if sess.UserID > 0 {
			return true, sess
		}
	}
	return false, nil
}

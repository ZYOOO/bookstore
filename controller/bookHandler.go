package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

// GetBooks 获取全部图书
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetBooks()
	//这里是不是可以不用模板,直接用w.Write然后前端对data进行处理, 嫁鸡随鸡嫁狗随狗吧
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//同时跳转页面
	t.Execute(w, books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bID := r.FormValue("bID")
	dao.DeleteBook(bID)
	GetBooks(w, r)
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
	GetBooks(w, r)
}

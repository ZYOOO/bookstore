package test

import (
	"bookstore/dao"
	"bookstore/model"
	"fmt"
	"testing"
)

func TestGetBooks(t *testing.T) {
	books, _ := dao.GetBooks()
	//遍历得到每一本图书
	for k, v := range books {
		fmt.Printf("第%v本图书的信息是：%v \n", k+1, v)
	}
}

func TestAddBooks(t *testing.T) {
	book := &model.Book{
		Title:   "is",
		Author:  "asd",
		Price:   0.99,
		Sales:   12,
		Stock:   122,
		ImgPath: "asdsdasdad",
	}
	dao.AddBook(book)
}

func TestGetBookByBID(t *testing.T) {
	book, _ := dao.GetBookByBID("12")
	fmt.Println(book)
}

func TestUpdateBook(t *testing.T) {
	book := &model.Book{
		BID:     22,
		Title:   "is",
		Author:  "asd",
		Price:   0.99,
		Sales:   12,
		Stock:   122,
		ImgPath: "asdsdasdad",
	}
	dao.UpdateBook(book)
}

func TestGetPageBooks(t *testing.T) {
	page, _ := dao.GetPageBooksByPrice("1", "10", "30")
	fmt.Println(page.TotalPage)
	fmt.Println(page.TotalRecord)
	fmt.Println(page.PageSize)
	fmt.Println(page.PageNo)
	for k, v := range page.Books {
		fmt.Println(k, v)
	}
}

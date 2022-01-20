package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestGetBooks(t *testing.T) {
	books, _ := GetBooks()
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
	AddBook(book)
}

func TestGetBookByBID(t *testing.T) {
	book, _ := GetBookByBID("12")
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
	UpdateBook(book)
}

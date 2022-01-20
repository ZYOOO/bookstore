package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

// GetBooks 获取所有图书
func GetBooks() ([]*model.Book, error) {
	//写sql语句
	sqlStr := "select bid,title,author,price,sales,stock,img_path from books"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		rows.Scan(&book.BID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}
	return books, nil
}

// AddBook 添加一本图书
func AddBook(b *model.Book) error {
	sql := "insert into books(title,author,price,sales,stock,img_path) values (?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBook 删除一本图书
func DeleteBook(bID string) error {
	sql := "delete from books where bid = ?"
	_, err := utils.Db.Exec(sql, bID)
	if err != nil {
		return err
	}
	return nil
}

// GetBookByBID 获取一本图书用于修改
func GetBookByBID(bID string) (*model.Book, error) {
	sql := "select bid,title,author,price,sales,stock,img_path from books where bid = ?"
	row := utils.Db.QueryRow(sql, bID)
	book := &model.Book{}
	row.Scan(&book.BID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

// UpdateBook 更新修改后的图书
func UpdateBook(b *model.Book) error {
	sql := "update books set title = ?,author = ?,price = ?,sales = ?,stock = ? where bid = ?"
	_, err := utils.Db.Exec(sql, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.BID)
	if err != nil {
		fmt.Println("UpdateBook error")
		return err
	}
	return nil
}

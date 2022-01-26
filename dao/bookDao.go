package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"strconv"
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

// GetPageBooks 获取分页的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取图书的总记录数
	sql := "select count(*) from books"
	var totalRecord int64
	row := utils.Db.QueryRow(sql)
	row.Scan(&totalRecord)
	//设置每页只显示四条记录
	var pageSize int64 = 4
	var totalPage int64
	//获取总页数
	if totalRecord%pageSize == 0 {
		totalPage = totalRecord / pageSize
	} else {
		totalPage = totalRecord/pageSize + 1
	}
	var books []*model.Book
	sql2 := "select bid,title,author,price,sales,stock,img_path from books limit ?,?"
	rows, _ := utils.Db.Query(sql2, (iPageNo-1)*pageSize, pageSize)
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.BID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPage:   totalPage,
		TotalRecord: totalRecord,
	}
	return page, nil
}

func GetPageBooksByPrice(pageNo, min, max string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取图书的总记录数
	sql := "select count(*) from books where price between  ? and ?"
	var totalRecord int64
	row := utils.Db.QueryRow(sql, min, max)
	row.Scan(&totalRecord)
	//设置每页只显示四条记录
	var pageSize int64 = 4
	var totalPage int64
	//获取总页数
	if totalRecord%pageSize == 0 {
		totalPage = totalRecord / pageSize
	} else {
		totalPage = totalRecord/pageSize + 1
	}
	var books []*model.Book
	sql2 := "select bid,title,author,price,sales,stock,img_path from books where price between  ? and ? limit ?,?"
	rows, _ := utils.Db.Query(sql2, min, max, (iPageNo-1)*pageSize, pageSize)
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.BID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPage:   totalPage,
		TotalRecord: totalRecord,
	}
	return page, nil
}

//UpdateBookCount 根据购物项中的相关信息更新购物项中图书的数量和金额小计
func UpdateBookCount(cartItem *model.CartItem) error {
	//写sql语句
	sql := "update cart_items set count = ? , amount = ? where book_id = ? and cart_id = ?"
	//执行
	_, err := utils.Db.Exec(sql, cartItem.Count, cartItem.GetAmount(), cartItem.Book.BID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

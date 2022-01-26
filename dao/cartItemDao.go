package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

func AddCartItem(cartItem *model.CartItem) error {
	sql := "insert into cart_items (count,amount,book_id,cart_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sql, cartItem.Count, cartItem.GetAmount(), cartItem.Book.BID, cartItem.CartID)
	if err != nil {
		fmt.Println("AddCartItem error:", err)
		return err
	}
	return nil
}

func GetCartItemByBookIDAndCartID(bID, cartID string) (*model.CartItem, error) {
	sql := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	stmt, err := utils.Db.Prepare(sql)
	if err != nil {
		fmt.Println("GetCartItemByBookID", err)
		return nil, err
	}
	row := stmt.QueryRow(bID, cartID)
	cartItem := &model.CartItem{}
	err2 := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err2 != nil {
		return nil, err2
	}
	cartItem.Book, _ = GetBookByBID(bID)
	return cartItem, nil
}

func GetCartItemByCartID(cartID string) ([]*model.CartItem, error) {
	sql := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	stmt, err := utils.Db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rows, err2 := stmt.Query(cartID)
	if err2 != nil {
		fmt.Println("GetCartItemByCartID", err2)
		return nil, err2
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		var bID string
		cartItem := &model.CartItem{}
		err3 := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bID, &cartItem.CartID)
		if err3 != nil {
			return nil, err3
		}
		book, _ := GetBookByBID(bID)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

func UpdateItem(cartItem *model.CartItem) error {
	sql := "update cart_items set count = ?, amount = ? where id = ?"
	stmt, err := utils.Db.Prepare(sql)
	if err != nil {
		fmt.Println("Update", err)
		return err
	}
	_, err2 := stmt.Exec(cartItem.Count, cartItem.Amount, cartItem.CartItemID)
	if err2 != nil {
		fmt.Println("UpdateItem Exec", err2)
		return err2
	}
	return nil
}

func DeleteCartItemsByCartID(cartID string) error {
	sql := "delete from cart_items where cart_id = ?"
	_, err := utils.Db.Exec(sql, cartID)
	if err != nil {
		fmt.Println("clear cart", err)
		return err
	}
	return nil
}

func DeleteCartItemByID(cartItemID string) error {
	sql := "delete from cart_items where id = ?"
	_, err := utils.Db.Exec(sql, cartItemID)
	if err != nil {
		fmt.Println("clear cart", err)
		return err
	}
	return nil
}

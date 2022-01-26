package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

func AddCart(cart *model.Cart) error {
	sql := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sql, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		fmt.Println("AddCart error:", err)
		return err
	}
	for _, cartItem := range cart.CartItems {
		AddCartItem(cartItem)
	}
	return nil
}

func GetCartByUserID(userID int64) (*model.Cart, error) {
	sql := "select id,total_count,total_amount,user_id from carts where user_id = ?"
	stmt, err := utils.Db.Prepare(sql)
	if err != nil {
		fmt.Println("GetCartByUserID:", err)
		return nil, err
	}
	row := stmt.QueryRow(userID)
	cart := &model.Cart{}
	err2 := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err2 != nil {
		fmt.Println("GetCartByUserID:", err2)
		return nil, err2
	}
	cart.CartItems, _ = GetCartItemByCartID(cart.CartID)
	return cart, nil
}

func UpdateCart(cart *model.Cart) error {
	sql := "update carts set total_count = ?,total_amount = ? where id = ?"
	_, err := utils.Db.Exec(sql, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		fmt.Println("UpdateCart", err)
		return err
	}
	return nil
}

func DeleteCartByCartID(cartID string) error {
	//先删除购物项
	DeleteCartItemsByCartID(cartID)
	//删除购物车
	sql := "delete from carts where id = ?"
	_, err := utils.Db.Exec(sql, cartID)
	if err != nil {
		fmt.Println("clear cart", err)
		return err
	}
	return nil
}

package test

import (
	"bookstore/dao"
	"bookstore/model"
	"fmt"
	"testing"
)

func TestAddCart(t *testing.T) {
	//设置要买的第一本书
	book := &model.Book{
		BID:   1,
		Price: 211.0,
	}
	//设置要买的第二本书
	book2 := &model.Book{
		BID:   2,
		Price: 23.00,
	}
	//创建一个购物项切片
	var cartItems []*model.CartItem
	//创建两个购物项
	cartItem := &model.CartItem{
		Book:   book,
		Count:  10,
		CartID: "66668888",
	}
	cartItems = append(cartItems, cartItem)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "66668888",
	}
	cartItems = append(cartItems, cartItem2)
	//创建购物车
	cart := &model.Cart{
		CartID:    "66668888",
		CartItems: cartItems,
		UserID:    5,
	}
	//将购物车插入到数据库中
	dao.AddCart(cart)
}

func TestGetCartItemByCartID(t *testing.T) {
	cartItems, _ := dao.GetCartItemByCartID("66668888")
	for i, v := range cartItems {
		fmt.Println(i, v)
	}
}

func TestGetCartByUserID(t *testing.T) {
	cart, _ := dao.GetCartByUserID(5)
	fmt.Println(cart.TotalCount)
	fmt.Println(cart.TotalAmount)
	for i, v := range cart.CartItems {
		fmt.Println(i, v)
	}
}

func TestGetCartItemByBookIDAndCartID(t *testing.T) {
	cartItem, _ := dao.GetCartItemByBookIDAndCartID("1", "b28e0d37-65dd-49de-4eaa-291d08f7e2c9")
	fmt.Println(cartItem)
}

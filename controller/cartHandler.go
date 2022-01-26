package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

func AddBookToCart(w http.ResponseWriter, r *http.Request) {
	//先判断是否登陆
	flag, session := IsLogin(r)
	if !flag {
		w.Write([]byte("请先登录!"))
		return
	}
	bID := r.FormValue("bID")
	iBID, _ := strconv.ParseInt(bID, 10, 64)
	book, _ := dao.GetBookByBID(bID)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil {
		//该用户已经有购物车
		//这本书是否已经在购物车里面了
		cartItem, _ := dao.GetCartItemByBookIDAndCartID(bID, cart.CartID)
		if cartItem != nil {
			//已经有这本书了,直接改count就行
			for _, v := range cart.CartItems {
				if iBID == v.Book.BID {
					v.Count += 1
					v.Amount = v.GetAmount()
					dao.UpdateItem(v)
					break
				}
			}
		} else {
			//还没此图书,得创建一个新购物项
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cart.CartID,
			}
			cart.CartItems = append(cart.CartItems, cartItem)
			dao.AddCartItem(cartItem)
		}
		dao.UpdateCart(cart)
	} else {
		//还未有购物车
		newCart := &model.Cart{
			CartID: utils.CreateUUID(),
			UserID: userID,
		}
		var cartItems []*model.CartItem
		cartItem := &model.CartItem{
			Book:   book,
			Count:  1,
			CartID: newCart.CartID,
		}
		cartItems = append(cartItems, cartItem)
		newCart.CartItems = cartItems
		dao.AddCart(newCart)
	}
	w.Write([]byte(book.Title))
}

func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	//先判断是否登陆
	flag, session := IsLogin(r)
	if !flag {
		w.Write([]byte("请先登录!"))
		return
	}
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
	if cart != nil {
		cart.Username = session.Username
		session.Cart = cart
		t.Execute(w, session)
	} else {
		//该用户还没有有购物车
		t.Execute(w, session)
	}
}

func ClearCart(w http.ResponseWriter, r *http.Request) {
	cartID := r.FormValue("cartID")
	dao.DeleteCartByCartID(cartID)
	GetCartInfo(w, r)
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要删除的购物项的id
	cartItemID := r.FormValue("cartItemID")
	//将购物项的id转换为int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//获取session
	_, session := IsLogin(r)
	//获取用户的id
	userID := session.UserID
	//获取该用户的购物车
	cart, _ := dao.GetCartByUserID(userID)
	//获取购物车中的所有的购物项
	cartItems := cart.CartItems
	//遍历得到每一个购物项
	for k, v := range cartItems {
		//寻找要删除的购物项
		if v.CartItemID == iCartItemID {
			//这个就是我们要删除的购物项
			//将当前购物项从切片中移出
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			//将删除购物项之后的切片再次赋给购物车中的切片
			cart.CartItems = cartItems
			//将当前购物项从数据库中删除
			dao.DeleteCartItemByID(cartItemID)
		}
	}
	//更新购物车中的图书的总数量和总金额
	dao.UpdateCart(cart)
	//调用获取购物项信息的函数再次查询购物车信息
	GetCartInfo(w, r)
}

//UpdateCartItem 更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要更新的购物项的id
	cartItemID := r.FormValue("cartItemId")
	//将购物项的id转换为int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//获取用户输入的图书的数量
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	//获取session
	_, session := IsLogin(r)
	//获取用户的id
	userID := session.UserID
	//获取该用户的购物车
	cart, _ := dao.GetCartByUserID(userID)
	//获取购物车中的所有的购物项
	cartItems := cart.CartItems
	//遍历得到每一个购物项
	for _, v := range cartItems {
		//寻找要更新的购物项
		if v.CartItemID == iCartItemID {
			//这个就是我们要更新的购物项
			//将当前购物项中的图书的数量设置为用户输入的值
			v.Count = iBookCount
			//更新数据库中该购物项的图书的数量和金额小计
			dao.UpdateBookCount(v)
		}
	}
	//更新购物车中的图书的总数量和总金额
	dao.UpdateCart(cart)
	//调用获取购物项信息的函数再次查询购物车信息
	cart, _ = dao.GetCartByUserID(userID)
	// GetCartInfo(w, r)
	//获取购物车中图书的总数量
	totalCount := cart.TotalCount
	//获取购物车中图书的总金额
	totalAmount := cart.TotalAmount
	var amount float64
	//获取购物车中更新的购物项中的金额小计
	cIs := cart.CartItems
	for _, v := range cIs {
		if iCartItemID == v.CartItemID {
			//这个就是我们寻找的购物项，此时获取当前购物项中的金额小计
			amount = v.Amount
		}
	}
	//创建Data结构
	data := model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}
	//将data转换为json字符串
	json, _ := json.Marshal(data)
	//响应到浏览器
	w.Write(json)
}

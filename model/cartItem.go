package model

// CartItem 购物车里面的购物项
type CartItem struct {
	CartItemID int64   //购物项的id
	Book       *Book   //购物项中的图书信息
	Count      int64   //购物项中图书的数量
	Amount     float64 //购物项之间的金额小计,直接通过计算得到就行了
	CartID     string
}

func (cartItem *CartItem) GetAmount() float64 {
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price
}

package model

type Cart struct {
	CartID      string
	CartItems   []*CartItem
	TotalCount  int64
	TotalAmount float64
	UserID      int64
	Username    string
}

func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range cart.CartItems {
		totalCount += v.Count
	}
	return totalCount
}

func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range cart.CartItems {
		totalAmount += v.GetAmount()
	}
	return totalAmount
}

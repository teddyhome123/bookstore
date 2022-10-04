package model

import _"fmt"
//購物項
type CartItem struct {
	CartItemID int //購物項ID
	Book *Book //圖書訊息
	Count int //圖書數量
	Amount int	//圖書金額小結 通過計算得到
	CartID string //當前購物項屬於哪個購物車ID
}

//獲取購物項金額小結
func (cartItem *CartItem) GetAmount() int {
	//獲取圖書價格
	price := cartItem.Book.Price
	//fmt.Println(cartItem.Count)
	return cartItem.Count * price
}
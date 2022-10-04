package model

import (
	//"time"
)

//訂單結構
type Order struct {
	OrderID string //訂單ID 唯一
	CreateTime string //下單時間
	TotalCount int //訂單總數
	TotalAmount int //訂單總金額
	State int //訂單狀態 0未發貨 1發貨 2交易完成
	UserID int //訂單所屬的用戶
}

//NoSend 未發貨
func (order *Order) NoSend() bool {
	return order.State == 0
}
//SendComplate已發貨
func (order *Order) SendComplate() bool {
	return order.State == 1
}
//Complate交易完成
func (order *Order) Complate() bool {
	return order.State == 2
}
package model

//Seesion結構
type Session struct {
	SessionID string  //透過UUID生成
	UserName string 
	UserID int 
	Cart *Cart
	Order *Order
	Orders []*Order
}


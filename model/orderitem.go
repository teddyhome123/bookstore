package model

//訂單項
type OrderItem struct {
	OrderItemID int //訂單項ID
	Count int //訂單項圖書數量
	Amount int //訂單項圖書價錢
	Title string //訂單項圖書書名
	Author string //訂單項圖書作者
	Price int //訂單項圖書的價錢
	ImgPath string //訂單圖書封面
	OrderID string //訂單項所屬的訂單
}
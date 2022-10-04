package controller
import (
	"bookstore/model"
	"bookstore/dao"
	"bookstore/utils"
	"net/http"
	"time"
	"html/template"
)

//checkout 去結帳
func Checkout(w http.ResponseWriter, r *http.Request) {
	//購取購物車的內容 session
	_, session := dao.IsLogin(r)
	userID := session.UserID
	//獲取購物車
	cart, _ := dao.GetCartByUserId(userID)
	//生成訂單號
	orderID := utils.CreateUUID()
	//生成訂單時間
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	//創建Order
	order := &model.Order{
		OrderID : orderID,
		CreateTime : timeStr,
		TotalCount : cart.TotalCount,
		TotalAmount : cart.TotalAmount,
		State : 0,
		UserID : userID,
	}
	//將訂單保存到資料庫
	dao.AddOrder(order)
	//保存訂單項
	//獲取購物車所有購物項
	cartItems := cart.CartItems
	for _, v := range cartItems {
		//創建訂單項
		orderItem := &model.OrderItem{
			Count : v.Count,
			Amount : v.Amount,
			Title : v.Book.Title,
			Author : v.Book.Author,
			Price : v.Book.Price,
			ImgPath : v.Book.ImgPath,
			OrderID : orderID,
		}
		//保存購物項到資料庫
		dao.AddOrderItem(orderItem)
		//更新當前購物項中圖書的庫存和銷量
		book := v.Book
		book.Sales = book.Sales + v.Count
		book.Stock = book.Stock - v.Count
		//更新圖書訊息
		dao.UpdateBook(book)
	}
	//清空購物車
	dao.DeleteCartByCartID(cart.CartID)
	//訂單號設置到session中
	session.Order = order
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, session)
}

//GetOrders 獲取所有訂單
func GetOrders(w http.ResponseWriter, r *http.Request) {
	//獲取所有訂單
	orders, _ := dao.GetOrders()

	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

//GetOrderInfo 獲取訂單詳情
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	//獲取訂單號
	orderID := r.FormValue("orderid")
	//調用方法
	orderItems, _ :=dao.GetOrderItemsByOrderID(orderID)

	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

//GetMtOrders 獲取我的訂單
func GetMyOrder(w http.ResponseWriter, r *http.Request) {
	//獲取用戶ID
	_, session := dao.IsLogin(r)
	//獲取用戶ID
	userID := session.UserID
	orders, _ := dao.GetMyOrder(userID)
	//將訂單設定到session
	session.Orders = orders
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, session)
}

//SendOrder 發貨
func SendOrder(w http.ResponseWriter, r *http.Request) {
	//獲取訂單號
	orderID := r.FormValue("orderId")
	//調用
	dao.UpdateOrderState(orderID, 1)
	//調用更新頁面
	GetOrders(w, r)
}

//TakeOrder收貨
func TakeOrder(w http.ResponseWriter, r *http.Request) {
	//獲取訂單號
	orderID := r.FormValue("orderId")
	//調用
	dao.UpdateOrderState(orderID, 2)
	//調用更新頁面
	GetMyOrder(w, r)
}
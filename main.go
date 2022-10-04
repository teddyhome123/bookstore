package main

import (
	"net/http"
	_"text/template"
	"bookstore/controller"
)




func main() {
	//處理靜態資源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	//去首頁
	http.HandleFunc("/main", controller.GetPageBooksByPrice)

	//去登入頁面
	http.HandleFunc("/login", controller.Login)
	//去登出頁面
	http.HandleFunc("/logout", controller.Logout)
	
	//去註冊頁面
	http.HandleFunc("/regist", controller.Regist)

	//通過AJAX請求驗證用戶是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)

	//獲取所有圖書
	//http.HandleFunc("/getBooks", controller.GetBooks)
	//獲取帶分頁的圖書
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	//新增圖書
	//http.HandleFunc("/addBook", controller.AddBook)

	//刪除圖書
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//去更新圖書頁面
	http.HandleFunc("/toUpdataBook", controller.ToUpdataBook)
	//更新或添加圖書
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)
	//查詢書本價格
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	//添加圖書到購物車
	http.HandleFunc("/addBookCart", controller.AddBookCart)
	//獲取購物車訊息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	//清空購物車訊息
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	//刪除購物項
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	//更新購物項
	http.HandleFunc("/updatCountItem", controller.UpdatCountItem)
	//去結帳
	http.HandleFunc("/checkout", controller.Checkout)
	//獲取所有訂單
	http.HandleFunc("/getOrders", controller.GetOrders)
	//獲取訂單詳情
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	//獲取我的訂單
	http.HandleFunc("/getMyOrder", controller.GetMyOrder)
	//發貨功能
	http.HandleFunc("/sendOrder", controller.SendOrder)
	//確認收貨
	http.HandleFunc("/takeOrder", controller.TakeOrder)




	http.ListenAndServe(":8080", nil)
}
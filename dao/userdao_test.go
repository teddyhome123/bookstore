package dao

import (
	"testing"
	"fmt"
	"bookstore/model"
	//"time"
)

func TestMain(m *testing.M) {
	fmt.Println("測試book中的方法")
	m.Run()
}

func TestUser(t *testing.T) {
	//fmt.Println("測試")
	//t.Run("驗證用戶名或密碼", testLogin)
	//t.Run("驗證用戶名", testRegist)
	//t.Run("註冊", testSave)
	
}

func testLogin(t *testing.T){
	user, _ := CheckUserNameAndPassword("admin", "123456",)
	fmt.Println("獲取到的用戶:", user)
}

func testRegist(t *testing.T){
	user, _ := CheckUserName("admin")
	fmt.Println("獲取到的用戶:", user)
}

func testSave(t *testing.T){
	SaveUser("admin2", "123456", "adad@dasdas")
}

func TestBook(t *testing.T){
	fmt.Println("測試bookDao的相關函數")
	//t.Run("測試獲取圖書", testGetBooks)
	//t.Run("測試圖書", testAddBook)
	//t.Run("刪除圖書", testDeleteBook)
	//t.Run("獲取圖書", testGetBook)
	//t.Run("更新圖書", testUpdateBook)
	//t.Run("測試分頁", testGetPageBooks)
	//t.Run("測試分頁", testGetPageBooksByPrice)
	
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for i, v := range books {
		fmt.Printf("第%v本書的訊息是:, %v\n", i+1, v)
	}
}

func testAddBook(t *testing.T) {
	book := &model.Book {
		Title : "三國演義",
		Author : "羅貫中",
		Price : 300,
		Sales : 100,
		Stock : 100,
		ImgPath : "/static/img/default.jpg",
	}
	AddBook(book)
}

func testDeleteBook(t *testing.T) {
	//調用刪除
	DeleteBook("36")
}

func testGetBook(t *testing.T) {
	//調用
	book, _ := GetBookByID("1")
	fmt.Println(book)
}

func testUpdateBook(t *testing.T) {
	//調用更新
	book := &model.Book {
		ID : 40,
		Title : "紅樓夢",
		Author : "曹雪芹",
		Price : 450,
		Sales : 100,
		Stock : 100,
		ImgPath : "static/img/default.jpg",
	}
	UpdateBook(book)
}

func testGetPageBooks(t *testing.T) {
	page, _ := GetPageBooks("9")
	fmt.Println("當前頁:",page.PageNo)
	fmt.Println("總頁數:",page.TotalPageNo)
	fmt.Println("總筆數",page.TotalRecrd)
	fmt.Println("頁面下的圖書:")
	for _, v := range page.Books{
		fmt.Println("圖書有:", v)
	}

}


func testGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("1", "100", "300")
	fmt.Println("當前頁:",page.PageNo)
	fmt.Println("總頁數:",page.TotalPageNo)
	fmt.Println("總筆數",page.TotalRecrd)
	fmt.Println("頁面下的圖書:")
	for _, v := range page.Books{
		fmt.Println("圖書有:", v)
	}

}

func TestSession(t *testing.T) {
	fmt.Println("測試Session函數")
	//t.Run("測試添加Session", testAddSession)
	//t.Run("測試刪除", testDeleteSession)
	//t.Run("測試獲取", testGetSession)
}

func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID : "12345679845",
		UserName : "測試一",
		UserID : 6,
	}
	AddSession(sess)
}

func testDeleteSession(t *testing.T) {
	DeleteSession("12345679845")
}

func testGetSession(t *testing.T) {
	sess := GetSession("ee01eb16-f94c-4ea1-573a-363153b4bcab")
	fmt.Println(sess)
}

func TestCart(t *testing.T) {
	fmt.Println("測試購物車相關函數")
	//t.Run("測試添加購物車", testAddCart)
	//t.Run("根據圖書ID獲取購物項", testGetCartItemByBookId)
	//t.Run("根據購物車ID獲取所有購物項", testGetCartItemByCart)
	//t.Run("根據用戶ID獲取購物車", testGetCartByUserID)
	//t.Run("根據圖書和購物車的ID更新購物項的數量", testUpdateBookItemCount)
	//t.Run("根據購物車ID刪除購物車", testDeleteCartByCartID)
	//t.Run("刪除購物項", testCartItemID)
}

func testAddCart(t *testing.T) {
	//測試購物車 
	book := &model.Book {
		ID : 1,
		Price : 490,
	}
	book2 := &model.Book {
		ID : 2,
		Price : 435,
	}
	//購物項切片
	var cartItems []*model.CartItem
	//購物項
	cartItem := &model.CartItem {
		Book : book,
		Count : 2,
		CartID : "1234567",
	}
	cartItems = append(cartItems, cartItem)
	cartItem2 := &model.CartItem {
		Book : book2,
		Count : 3,
		CartID : "1234567",
	}
	cartItems = append(cartItems, cartItem2)

	cart := &model.Cart {
		CartID : "1234567",
		CartItems : cartItems,
		UserID : 4,
	}
	AddCart(cart)
}

func testGetCartItemByBookId(t *testing.T) {
	book, _ := GetCartItemByBookIDAndCartID("2", "1234567")
	fmt.Println("圖書購物項訊息", book)
}

func testGetCartItemByCart(t *testing.T) {
	cartItems, _ := GetCartItemByCartID("1234567")
	for k, v := range cartItems {
		fmt.Printf("第%v個購物項是: %v \n", k+1, v)
	}																								
}

func testGetCartByUserID(t *testing.T) {
	usercart, _ := GetCartByUserId(4)
	fmt.Println("購物車:", usercart)
}

func testUpdateBookItemCount(t *testing.T) {
	//UpdateBookItemCount(11, 1, 120,  "1234567")
}

func testDeleteCartByCartID(t *testing.T) { 
	DeleteCartByCartID("612db8d4-2549-4dd3-5903-810c61fcbbc9")
}

func testCartItemID(t *testing.T) { 
	DeleteCartItemByID("42")
}

func TestOrder(t *testing.T) {
	fmt.Println("測試訂單")
	//t.Run("測試添加訂單", testAddorder)
	//t.Run("取得訂單", testGetOrders)
	//t.Run("取得訂單", testGetOrderItems)
	t.Run("取得我得訂單", testGetMyOrders)
	t.Run("更新訂單狀態", testUpdateOrderState)
}

func testAddorder(t *testing.T) {
	//創建訂單
	order := &model.Order{
		OrderID : "1234567",
		//CreateTime : time.Now(),
		TotalCount : 2,
		TotalAmount : 1000,
		State : 0,
		UserID : 1,
	}
	//訂單項
	orderItem := &model.OrderItem {
		Count : 2,
		Amount : 1000,
		Title : "三國演義",
		Author : "羅貫中",
		Price : 500,
		ImgPath : "static/img/default.jpg",
		OrderID : "1234567",
	}	
		AddOrder(order)
		AddOrderItem(orderItem)
}

func testGetOrders(t *testing.T) {
	orders, _ := GetOrders()
	for _, v := range orders {
		fmt.Println(v)
	}
}

func testGetOrderItems(t *testing.T) {
	orders, _ := GetOrderItemsByOrderID("998c7149-6a55-4e43-52ca-6bedd1b6bd6a")
	for _, v := range orders {
		fmt.Println(v)
	}
}

func testGetMyOrders(t *testing.T) { 
	orders, _ := GetMyOrder(1)
	for _, v := range orders {
		fmt.Println(v)
	}
}

func testUpdateOrderState(t *testing.T) {
	UpdateOrderState("02cf79d2-8d84-4f61-4674-f34b21a7d71e", 1)
}

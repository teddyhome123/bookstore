package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

// //去首頁
// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	//獲取頁碼
// 	pageNo := r.FormValue("pageNo")
// 	if pageNo == "" {
// 		pageNo = "1"
// 	}

// 	//調用bookdao帶分頁的函數
// 	page, _ := dao.GetPageBooks(pageNo)

// 	//解析模板
// 	t := template.Must(template.ParseFiles("views/index.html"))
// 	t.Execute(w, page)
// }

//GetpageBooks 獲取帶有分頁的圖書
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//獲取頁碼
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}

	//調用bookdao帶分頁的函數
	page, _ := dao.GetPageBooks(pageNo)
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//執行
	t.Execute(w, page)
}


//(首頁) GetPageBooksByPrice 獲取帶有分頁和價格範圍的圖書
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//獲取頁碼
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//獲取價格範圍
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		//調用bookdao帶分頁的函數
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		//調用bookdao帶分頁的函數
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)

		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}

	//調用IsLigin()判斷有無登入
	flag, session := dao.IsLogin(r)
		if flag {
			page.IsLogin = true
			page.Username = session.UserName
		}

	//解析模板文件
	t := template.Must(template.ParseFiles("views/index.html"))
	//執行
	t.Execute(w, page)
}

//GetBooks 獲取所有圖書
// func GetBooks(w http.ResponseWriter, r *http.Request) {
// 	//調用bookdao函數
// 	books, _ := dao.GetBooks()
// 	//解析模板文件
// 	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
// 	//執行
// 	t.Execute(w, books)
// }

//AddBook 添加圖書
// func AddBook(w http.ResponseWriter, r *http.Request) {
// 	//從前端獲取圖書訊息
// 	title := r.PostFormValue("title")
// 	author := r.PostFormValue("author")
// 	price := r.PostFormValue("price")
// 	sales := r.PostFormValue("sales")
// 	stock := r.PostFormValue("stock")
// 	//將回傳過來是string類型 進行轉換
// 	intPrice, _ := strconv.ParseInt(price, 10, 0)
// 	intSales, _ := strconv.ParseInt(sales, 10, 0)
// 	intstock, _ := strconv.ParseInt(stock, 10, 0)
// 	//創建Book
// 	book := &model.Book {
// 		Title : title,
// 		Author : author,
// 		Price : int(intPrice),
// 		Sales : int(intSales),
// 		Stock : int(intstock),
// 		ImgPath : "static/img/default.jpg",
// 	}
// 	dao.AddBook(book)
// 	//返回頁面 再次查詢
// 	GetBooks(w, r)
// }

//Deletebook刪除圖書
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//從前端獲取圖書Id
	bookID := r.FormValue("bookID")
	//調用刪除bookdao中的函數
	dao.DeleteBook(bookID)
	//返回頁面 再次查詢
	GetPageBooks(w, r)
}

//ToUpdataBook去更新或添加圖書的頁面
func ToUpdataBook(w http.ResponseWriter, r *http.Request) {
	//從前端獲取圖書Id
	bookID := r.FormValue("bookID")
	//調用修改bookdao中的函數
	book, _ := dao.GetBookByID(bookID)
	if book.ID > 0 {
		//有ID更新圖書
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//執行
		t.Execute(w, book)
	} else {
		//沒有ID 添加圖書
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//執行
		t.Execute(w, "")
	}
}

//UpdateBook更新或添加圖書
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	//從前端獲取圖書訊息
	bookID := r.PostFormValue("bookID")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	//將回傳過來是string類型 進行轉換
	intbookID, _ := strconv.ParseInt(bookID, 10, 0)
	intPrice, _ := strconv.ParseInt(price, 10, 0)
	intSales, _ := strconv.ParseInt(sales, 10, 0)
	intstock, _ := strconv.ParseInt(stock, 10, 0)
	//創建Book
	book := &model.Book {
		ID : int(intbookID),
		Title : title,
		Author : author,
		Price : int(intPrice),
		Sales : int(intSales),
		Stock : int(intstock),
		ImgPath : "static/img/default.jpg",
	}
	if book.ID > 0 {
		//新增
		dao.UpdateBook(book)
	} else {
		//添加
		dao.AddBook(book)
	}
	
	//返回頁面 再次查詢
	GetPageBooks(w, r)
}


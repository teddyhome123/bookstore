package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"strconv"
)

//獲取資料庫所有的書籍
func GetBooks()([]*model.Book, error) {
	//SQL語法
	sqlStr := "select * from books"
	//執行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//給book中的字段賦值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//將book添加到books中
		books = append(books, book)
	}
	return books, nil
}

//向資料庫中添加圖書
func AddBook(b *model.Book) error {
	//SQL語法
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	//執行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//根據圖書ID刪除圖書
func DeleteBook(bookID string) error {
	//SQL語法
	sqlStr := "delete from books where id = ?"
	//執行
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

//GetBookById 根據圖書Id從資料庫查詢
func GetBookByID(bookID string) (*model.Book, error) {
	//SQL語法
	sqlStr := "select * from books where id = ?"
	//執行
	row := utils.Db.QueryRow(sqlStr, bookID)
	//創建一個Book
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)

	return book, nil
} 

//updateBook根據圖書ID修改圖書
func UpdateBook(b *model.Book) error {
	//SQL語法
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id = ?"
	//執行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return err
	}
	return nil
}

//GetPageBooks 獲取分頁的圖書訊息
func GetPageBooks(pageNo string) (*model.Page, error) {
	//將頁碼轉成int
	intPageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return nil, err
	}

	//獲取圖書總數
	sqlStr := "select count(*) from books"
	//紀錄總數
	var totalRecord int64
	//執行
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	//設置每頁只顯示四筆紀錄
	var pageSize int64
	pageSize = 4
	//獲取總頁數
	var totalPageNo int64
	if totalRecord % pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord / pageSize + 1
	}
	//獲取當前頁中的圖書
	sqlStr2 := "select * from books limit ?, ?"
	//執行
	var books []*model.Book
	rows, err := utils.Db.Query(sqlStr2, (intPageNo - 1) * pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
			book := &model.Book {}
			rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
			//將book添加到books
			books = append(books, book)
	}
	//創建page
	page := &model.Page {
		Books : books,
		PageNo : intPageNo,
		PageSize : pageSize,
		TotalPageNo : totalPageNo,
		TotalRecrd : totalRecord,
	}
	return page, nil
}

//GetPageBooks 獲取分頁和價格的圖書訊息 (首頁)
func GetPageBooksByPrice(pageNo string, minPrice string, maxPrice string) (*model.Page, error) {
	//將頁碼轉成int
	intPageNo, err := strconv.ParseInt(pageNo, 10, 64)
	if err != nil {
		return nil, err
	}

	//獲取圖書總數
	sqlStr := "select count(*) from books where price between ? and ?"
	//紀錄總數
	var totalRecord int64
	//執行
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)
	//設置每頁只顯示四筆紀錄
	var pageSize int64
	pageSize = 4
	//獲取總頁數
	var totalPageNo int64
	if totalRecord % pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord / pageSize + 1
	}
	//獲取當前頁中的圖書
	sqlStr2 := "SELECT * FROM books WHERE price BETWEEN ? AND ? ORDER BY price ASC LIMIT ?, ?"
	//執行
	var books []*model.Book
	rows, err := utils.Db.Query(sqlStr2, minPrice, maxPrice, (intPageNo - 1) * pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
			book := &model.Book {}
			rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
			//將book添加到books
			books = append(books, book)
	}
	//創建page
	page := &model.Page {
		Books : books,
		PageNo : intPageNo,
		PageSize : pageSize,
		TotalPageNo : totalPageNo,
		TotalRecrd : totalRecord,
	}
	return page, nil
}
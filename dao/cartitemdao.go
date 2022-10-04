package dao

import (
	"bookstore/utils"
	"bookstore/model"
)

//向購物項表中插入購物項
func AddCartItem(cartItem *model.CartItem) error {
	//SQL
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	//執行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(),
	cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

//GetCartItemByBookID 根據圖書ID獲取對應的購物項目 (檢查圖書有無在購物車內)
func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	//SQL
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	//執行
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)
	//cartItem實例
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	//根據圖書ID查詢圖書
	book, _ := GetBookByID(bookID)
	//將book設置到購物項中
	cartItem.Book = book 
	return cartItem, nil
}

//UpdateBookCount 根據圖書ID和購物車ID更新圖書數量
func UpdateBookItemCount(cartItem *model.CartItem) error {
	//var updateAmount *model.CartItem
	//SQL
	sqlStr := "update cart_items set count = ?, amount = ? where book_id = ? and cart_id = ?"
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil 
}

//GetCartItemByCartID 根據購物車ID獲取購物車中所有購物項
func GetCartItemByCartID(cartID string) ([]*model.CartItem, error) {
	//SQL
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	//執行
	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		var bookID string
		//cartItem實例
		cartItem := &model.CartItem{}
		err := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err != nil {
			return nil, err
		}
		//根據bookid查詢圖書訊息
		book, _ := GetBookByID(bookID)
		//book設置到購物項中
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//DeleteCartItemsByCartID 根據購物車ID刪除所有購物項
func DeleteCartItemsByCartID(cartID string) error {
	//SQL
	sqlStr := "delete from cart_items where cart_id = ?"
	_, err := utils.Db.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCartItemByID 根據購物項的ID刪除購物項 (單個)
func DeleteCartItemByID(cartItemID string) error {
	//SQL
	sqlStr := "delete from cart_items where id = ?"
	_, err := utils.Db.Exec(sqlStr, cartItemID)
	if err != nil {
		return err
	}
	return nil 
}

package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

//AddCart 向購物車表中插入購物車
func AddCart(cart *model.Cart) error {

	//Sql
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	//執行
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}

	//獲取購物車中所有購物項
	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		//保存購物項插入到資料庫
		AddCartItem(cartItem)
	}
	return nil
}

//根據用戶ID從資料庫查詢對應的購物車
func GetCartByUserId(userID int) (*model.Cart, error) {
	//SQL
	sqlStr := "select * from carts where user_id = ?"
	//執行
	row := utils.Db.QueryRow(sqlStr, userID)
	//購物車實例
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	//獲取購物車所有購物項
	cartItems, _ := GetCartItemByCartID(cart.CartID)
	//所有購物項設置到購物車中
	cart.CartItems = cartItems
	return cart, nil
}

//UpdateCart 更新購物車中的圖書總數量和金額
func UpdateCart(cart *model.Cart) error {
	//SQL
	sqlStr := "update carts set total_count = ?, total_amount = ? where id = ?"
	//執行
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil 
}

//DeleteCartByCartID 根據購物車ID刪除購物車
func DeleteCartByCartID(cartID string) error {
	//刪除購物車前先刪除所有購物項 (關聯表)
	err := DeleteCartItemsByCartID(cartID)
	if err != nil {
		return err 
	}
	//SQL
	sqlStr := "delete from carts where id = ?"
	_, err = utils.Db.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil 
}
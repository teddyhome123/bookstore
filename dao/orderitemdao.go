package dao

import (
	"bookstore/utils"
	"bookstore/model"
	
)


// AddOrderItem向資料庫中添加訂單項
func AddOrderItem(orderItem *model.OrderItem) error {
	//Sql
	sqlStr := "insert into order_items(count, amount, title, author, price, img_path, order_id) values (?,?,?,?,?,?,?)"
	//執行
	_, err := utils.Db.Exec(sqlStr, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price,
	orderItem.ImgPath, orderItem.OrderID)
	if err != nil {
		return err
	}
	return nil 
}

//GetOrderItemsByOrderID 根據訂單號獲取所有訂單項
func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	//SQL
	sqlStr := "select id, count, amount, title, author, price, img_path, order_id from order_items where order_id = ?"
	//執行
	rows, err := utils.Db.Query(sqlStr, orderID)
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.OrderItemID, &orderItem.Count, &orderItem.Amount, &orderItem.Title,
			 &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderID)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
package dao

import (
	"bookstore/utils"
	"bookstore/model"
)

//AddOrder 向資料庫中插入訂單
func AddOrder(order *model.Order) error {
	//SQL
	sqlStr := "insert into orders values(?,?,?,?,?,?)"
	//執行
	_, err := utils.Db.Exec(sqlStr, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil 
}

//GetOrders獲得資料庫所有訂單
func GetOrders() ([]*model.Order, error) {
	//SQL
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders"
	//執行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount,
			&order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}

//getMyOrder 獲取個人訂單訊息
func GetMyOrder(userID int) ([]*model.Order, error) {
	//SQL
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id = ?"
	//執行
	rows, err := utils.Db.Query(sqlStr, userID)
	if err != nil {
		return nil, err
	}
	//切片
	var orders []*model.Order
	for rows.Next() {
		//創建order結構體
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount,
		&order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil 
}

//UpdateOrderState 更新訂單狀態
func UpdateOrderState(orderID string, state int) error {
	//SQL
	sqlStr := "update orders set state = ? where id = ?"
	//執行
	_, err := utils.Db.Exec(sqlStr, state, orderID)
	if err != nil {
		return err
	}
	return nil 
}
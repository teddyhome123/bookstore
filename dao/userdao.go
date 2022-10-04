package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

//查詢用戶與用戶密碼
func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	//查詢
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	//執行只查詢一條
	row := utils.Db.QueryRow(sqlStr, username, password)
	//創建結構體實例
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}


//查詢用戶名
func CheckUserName(username string) (*model.User, error) {
	//查詢
	sqlStr := "select id,username,password,email from users where username = ?"
	//執行只查詢一條
	row := utils.Db.QueryRow(sqlStr, username)
	//創建結構體實例
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

//註冊 向資料庫中寫入用戶
func SaveUser(username string, password string, email string) error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//執行
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}
	return nil
}
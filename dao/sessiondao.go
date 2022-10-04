package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"net/http"
)

//AddSession 向資料庫中添加Session
func AddSession(sess *model.Session) error {
	//SQL語法
	sqlStr := "insert into sessions values(?,?,?)"
	//執行
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession 刪除
func DeleteSession(sessID string) error {
	//SQL語法
	sqlStr := "delete from sessions where session_id = ?"
	//執行
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

//GetSessionById 從數據庫中取出ID
func GetSession(sessId string) (*model.Session) {
	//SQL語法
	sqlStr := "select * from sessions where session_id = ?"
	//執行
	row := utils.Db.QueryRow(sqlStr, sessId)
	//創建Session實例
	sess := &model.Session{}
	//掃描資料庫中的值
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)

	return sess
}

//判斷用戶是否登入 false沒有登入 true登入
func IsLogin(r *http.Request) (bool, *model.Session) {
	//根據cookie的name獲取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//獲取cookie的value
		cookieValue := cookie.Value
		//根據cookieValue去資料庫查詢對應的session
		session := GetSession(cookieValue)

		if session.UserID > 0 { //資料庫有此Session
			//已經登入 設置page中的IsLogin和UserName
			return true, session		
		}
	}
	return false, nil
}
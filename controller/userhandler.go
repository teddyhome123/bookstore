package controller

import (
	"bookstore/utils"
	"bookstore/model"
	"net/http"
	"bookstore/dao"
	"html/template"
	_"fmt"
)

//Login處理用戶登入的函數
func Login(w http.ResponseWriter, r *http.Request){
	//判斷是否已經登入
	flag, _ := dao.IsLogin(r)
	if flag {
		//已經登入
		//返回首頁
		GetPageBooksByPrice(w, r)
	} else {
		//獲取用戶名和密碼
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		//調用userdao驗證用戶名和密碼的方法
		user, _ := dao.CheckUserNameAndPassword(username, password)
		//判斷有這個ID
		if user.ID > 0 {
			//用戶名密碼正確
			//生成UUID
			uuid := utils.CreateUUID()
			//創建一個Session 保存用戶訊息
			sess := &model.Session{
				SessionID : uuid,
				UserName : user.Username,
				UserID : user.ID,
			}
			//將Session保存到資料庫
			dao.AddSession(sess)
			//創建一個cookie 讓他與sess關聯
			cookie := http.Cookie {
				Name : "user",
				Value : uuid,
				HttpOnly : true,
			}
			//將cookie發送給瀏覽器
			http.SetCookie(w, &cookie)

			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)	
		} else {
			//用戶名密碼錯誤
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "用戶名或密碼錯誤")	
		}

	}

}

//處理用戶登出
func Logout(w http.ResponseWriter, r *http.Request) {
	//獲取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//獲取cookie的值
		cookieValue := cookie.Value
		//刪除資料庫的Session
		dao.DeleteSession(cookieValue)
		//設置cookie失效
		cookie.MaxAge = -1 
		http.SetCookie(w, cookie)
	}
	//返回首頁
	GetPageBooksByPrice(w, r)
}

//Regist處理用戶註冊的函數
func Regist(w http.ResponseWriter, r *http.Request){
	//獲取用戶名和密碼
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//調用userdao驗證用戶名和密碼的方法
	user, _ := dao.CheckUserName(username)
	//判斷有這個ID
	if user.ID > 0 {
		//用戶名存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用戶名已存在")	
	} else {
		//用戶名不存在
		//將用戶保存到資料庫
		dao.SaveUser(username, password, email)

		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")	
	}
}

//checkUserName通過發送AJAX驗證用戶可用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//先獲取用戶輸入的用戶名
	username := r.PostFormValue("username")
	//fmt.Println("username=", username)

	user, _ := dao.CheckUserName(username)
	//判斷有這個ID
	if user.ID > 0 {
		//用戶存在
		w.Write([]byte("用戶名已存在"))
	} else {
		//用戶不存在
		w.Write([]byte("<font style='color:green'>用戶名稱可使用</font>"))
	}
}
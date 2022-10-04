package controller
import (
	"net/http"
	"bookstore/model"
	"bookstore/dao"
	"bookstore/utils"
	_"fmt"
	"html/template"
	"strconv"
)

//添加圖書到購物車
func AddBookCart(w http.ResponseWriter, r *http.Request){
	//判斷是否登入(取得session)
	flag, session := dao.IsLogin(r)
	//已經登入
	if flag {
	//獲取圖書ID
	bookID := r.FormValue("bookId")
	//fmt.Println("bookId=", bookId)
	//根據圖書ID獲取圖書訊息
	book, _ := dao.GetBookByID(bookID)
	//獲取UserId
	userID := session.UserID
	//判斷資料庫中是否有當前用戶的購物車
	cart, _ := dao.GetCartByUserId(userID)
	if cart != nil {
		//當前用戶已經有購物車
		//判斷購物車中是否有當前圖書
		cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
		if cartItem != nil {
			//購物車中的購物項目已經有該圖書
			//只需將圖書+1
			//1. 獲取購物車切片中所有的購物項
			cts := cart.CartItems
			//2. 遍歷所有購物項
			for _, v := range cts {
				//3. 找到當前購物項
				if v.Book.ID == cartItem.Book.ID {
					//fmt.Println("原本的Book",v)
					//fmt.Println("查詢到的Book-----------------------",cartItem.Book)
					//將當前購物項中的圖書數量+1
					v.Count = v.Count + 1 
					//fmt.Println(ItemsAmount)
					//更新資料庫內購物項的數量
					dao.UpdateBookItemCount(v)
				}
			}
		} else {
			//fmt.Println("沒有該圖書")
			//沒有該圖書 需創建一個新的購物項並添加到資料庫中
			cartItem2 := &model.CartItem{
				Book : book,
				Count : 1,
				CartID : cart.CartID,
			}
			//將購物項添加到當前cart切片中
			cart.CartItems = append(cart.CartItems, cartItem2)
			//將新創建的購物項插入到資料庫中
			dao.AddCartItem(cartItem2)
		}
		//更新總購物車
		dao.UpdateCart(cart)
	} else {
		//當前用戶沒有購物車 需要創建購物車並添加
		//1.創建購物車
			//1.1生成購物車ID
		cartID := utils.CreateUUID()
		cart := &model.Cart{
			CartID : cartID,
			UserID : userID,
		}
		//2.創建購物車中的購物項
			//2.1聲明一個切片
		var cartItems []*model.CartItem
		cartItem := &model.CartItem{
			Book : book,
			Count : 1,
			CartID : cartID,
		}
			//2.2將購物項添加到切片中
		cartItems = append(cartItems, cartItem)
		//3.將切片設置到cart中
		cart.CartItems = cartItems
		//4.將購物車保存到資料庫
		dao.AddCart(cart)
	}
	w.Write([]byte("您剛剛將   " + book.Title + "   放入到購物車"))
	} else {
		//沒有登入
		w.Write([]byte("請先登入"))
	}
}

//GetCartInfo獲取購物車訊息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	//判斷是否登入(取得session)
	_, session := dao.IsLogin(r)
	//獲取用戶ID
	userID := session.UserID
	//根據用戶ID從資料庫獲取對應的購物車
	cart, _ := dao.GetCartByUserId(userID)
	if cart != nil {
		//設置用戶名
		cart.UserName = session.UserName
		session.Cart = cart
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	} else {
		//該用戶沒有購物車
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	}
}

//DeleteCart 清空購物車
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	//獲取要刪除的購物車ID
	cartID := r.FormValue("cartId")
	//清空購物車
	dao.DeleteCartByCartID(cartID)
	//調用GetCartInfo
	GetCartInfo(w, r)
}

//DeleteCartItem 刪除購物項
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	// //獲取session
	_, session := dao.IsLogin(r)
	// //獲取用戶ID
	userID := session.UserID
	// //獲取該用戶的購物車
	cart, _ := dao.GetCartByUserId(userID)
	//獲取購物車中的購物項
	cartItems := cart.CartItems
	//遍歷
	for k, v := range cartItems {
		//尋找要刪除的購物項
		if v.CartItemID == int(iCartItemID) {
		//將當前購物項從切片中移除
		cartItems = append(cartItems[:k], cartItems[k+1:]...)
		//將刪除購物購項後的切片給購物車
		cart.CartItems = cartItems
		//將當前購物項從資料庫刪除
		dao.DeleteCartItemByID(cartItemID)
		}
	}
	dao.UpdateCart(cart)
	//獲取購物車訊息 再次查詢
	GetCartInfo(w, r)
}

//UpdatCountItem 更新購物項
func UpdatCountItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	//獲取用戶輸入的圖書數量
	bookCount := r.FormValue("bookCount")

	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	// //獲取session
	_, session := dao.IsLogin(r)
	// //獲取用戶ID
	userID := session.UserID
	// //獲取該用戶的購物車
	cart, _ := dao.GetCartByUserId(userID)
	//獲取購物車中的購物項
	cartItems := cart.CartItems
	//遍歷
	for _, v := range cartItems {
		//尋找要更新的購物項
		if v.CartItemID == int(iCartItemID) {
			//要更新的購物項 將圖書數量設置為用戶設定的值
			v.Count = int(iBookCount)
			//更新資料庫中該購物項的值
			dao.UpdateBookItemCount(v)
		}
	}
	dao.UpdateCart(cart)
	//獲取購物車訊息 再次查詢
	GetCartInfo(w, r)
}


func maxProfit(prices []int) int {
    minvalue := prices[0]
    var profit = 0
    for i := 1; i < len(prices); i++ {   
        if prices[i] < minvalue {
            minvalue = prices[i]
        } else {
            newprice := prices[i] - minvalue
            profit := newprice
            if newprice > profit {
                profit = newprice
            }
        }
    }
    return profit
}
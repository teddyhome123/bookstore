package model

//import "fmt"

//購物車
type Cart struct {
	//購物車ID
	CartID string
	CartItems []*CartItem //購物車中所有的購物項 存放到切片
	TotalCount int //購物車總數量
	TotalAmount int //購物車總金額
	UserID int //屬於哪個用戶的購物車
	UserName string
}

//GetTotalCount獲取購物車中圖書總數量
func (c *Cart) GetTotalCount() int {
	var totalCount int
	//遍歷購物車中購物項CartItems
	for _, v := range c.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

//GetTotalCount獲取購物車中圖書總金額
func (c *Cart) GetTotalAmount() int {
	var totalAmount int
	//遍歷購物車中購物項CartItems
	for _, v := range c.CartItems {
		//fmt.Println(v)
		totalAmount = totalAmount + v.GetAmount()
	}
	//fmt.Println(totalAmount)
	return totalAmount
}


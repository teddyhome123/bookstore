package model

//Page結構
type Page struct {
	Books []*Book //每頁查詢出來的圖書存放的切片
	PageNo int64 //當前頁
	PageSize int64 //每頁顯示的圖書
	TotalPageNo int64 //總頁數 通過計算得到
	TotalRecrd int64 //總筆數
	MinPrice string 
	MaxPrice string
	IsLogin bool
	Username string
}

//IsHasPrev 判斷是否有上一頁
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
} 

//IsHasNext 判斷是否有下一頁
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

//GetPrevPageNo 獲取上一頁
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

//GetNextPageNo 獲取下一頁
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}
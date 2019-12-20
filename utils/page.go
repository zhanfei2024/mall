package utils

var (
	PageSize int = 10
)

type Page struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalPage  int         `json:"totalPage"`
	TotalCount int         `json:"totalCount"`
	FirstPage  bool        `json:"firstPage"`
	LastPage   bool        `json:"lastPage"`
	List       interface{} `json:"list"`
}

func Pagination(count int, page int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{Page: page, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: page == 1, LastPage: page == tp, List: list}
}

package pkgutils

import (
	"math"
)

type Page struct {
	PageNo     int         `json:"page_num"`
	PageSize   int         `json:"page_size"`
	TotalPage  int         `json:"total_page"`
	TotalCount int         `json:"total_count"`
	FirstPage  bool        `json:"first_page"`
	LastPage   bool        `json:"last_page"`
	List       interface{} `json:"list"`
}


//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page, pageSize int, nums int64, list interface{}) map[string]interface{} {
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(pageSize))) //page总数

	if page <= 0 {
		page = 1
	}

	nextpage := page + 1
	if page >= totalpages {
		nextpage = 0
	}

	newList := ToSliceIface(list)

	paginatorMap := make(map[string]interface{})
	paginatorMap["next_page"] = nextpage
	paginatorMap["currpage"] = page
	paginatorMap["list"] = newList
	return paginatorMap
}

func GetPageParam(page, size int32) (int32, int32) {
	offset := (page -1) * size
	limit := size
	return offset,limit
}
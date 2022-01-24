package model

type Page struct {
	Books       []*Book //每页查询出来的图书存放的切片
	PageNo      int64   //当前页
	PageSize    int64
	TotalPage   int64
	TotalRecord int64
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	Username    string
}

func (p *Page) HasPrev() bool {
	return p.PageNo > 1
}

func (p *Page) HasNext() bool {
	return p.PageNo < p.TotalPage
}

func (p *Page) GetPrevPageNo() int64 {
	if p.HasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

func (p *Page) GetNextPageNo() int64 {
	if p.HasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPage
	}
}

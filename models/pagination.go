package models

type IPagination interface {
	GetPage() int
	SetPage(page int)
	GetSize() int
	SetSize(size int)
	GetTotal() int64
	SetTotal(total int64)
}

type Pagination struct {
	Page  int   `form:"page" json:"page"`
	Size  int   `form:"size" json:"size"`
	Total int64 `json:"total"` // this field is used in response, not in query
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}

func (p *Pagination) SetPage(page int) {
	p.Page = page
}

func (p *Pagination) GetSize() int {
	if p.Size <= 0 {
		return 10
	}
	return p.Size
}

func (p *Pagination) SetSize(size int) {
	p.Size = size
}

func (p *Pagination) GetTotal() int64 {
	return p.Total
}

func (p *Pagination) SetTotal(total int64) {
	p.Total = total
}

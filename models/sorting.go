package models

type ISorting interface {
	GetBy() string
	SetBy(by string)
	GetAsc() bool
	SetAsc(asc bool)
}

type Sorting struct {
	By  string `form:"by"`
	Asc bool   `form:"asc"`
}

func (s *Sorting) GetBy() string {
	return s.By
}

func (s *Sorting) SetBy(by string) {
	s.By = by
}

func (s *Sorting) GetAsc() bool {
	return s.Asc
}

func (s *Sorting) SetAsc(asc bool) {
	s.Asc = asc
}

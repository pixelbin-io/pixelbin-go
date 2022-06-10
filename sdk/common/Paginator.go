package common

//Paginator holds paginator object
type Paginator struct {
	PageNo      float64
	HasNextPage bool
	NextID      string
	Next        func() (interface{}, error)
}

//NewPaginator returns paginator instance
func NewPaginator(pageType string) *Paginator {
	if pageType == "number" {
		return &Paginator{
			PageNo:      1,
			NextID:      "",
			HasNextPage: true,
		}
	}
	if pageType == "cursor" {
		return &Paginator{
			PageNo:      0,
			NextID:      "*",
			HasNextPage: true,
		}
	}
	return &Paginator{}
}

//HasNext return has_next value
func (p *Paginator) HasNext() bool {
	return p.HasNextPage
}

//SetPaginator sets paginator instance with either next_id or page_no
func (p *Paginator) SetPaginator(hasNext bool, pageNo int, nextID string) {
	p.HasNextPage = hasNext
	p.NextID = nextID
	p.PageNo = float64(pageNo)
}

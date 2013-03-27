package library

type Page struct {
	Count     int
	PageNum   int
	PageSize  int
	TotalPage int
	Links     []int
	PreLink   int
	NextLink  int
	BaseUrl   string
	Visable   bool
}

func (p *Page) Compute() {
	if p.Count > p.PageSize {

		if p.Count%p.PageSize == 0 {
			p.TotalPage = p.Count / p.PageSize
		} else {
			p.TotalPage = p.Count/p.PageSize + 1
		}

		if p.PageNum > 1 {
			p.PreLink = p.PageNum - 1
		} else {
			p.PreLink = 1
		}

		if p.PageNum < p.TotalPage {
			p.NextLink = p.PageNum + 1
		} else {
			p.NextLink = p.TotalPage
		}

		if p.TotalPage > 5 {
			index := 1
			if p.PageNum > 3 && p.TotalPage-p.PageNum > 1 {
				index = p.PageNum - 2
			}
			if p.PageNum > 3 && p.TotalPage-p.PageNum < 2 {
				index = p.TotalPage - 5 + 1
			}

			for i := 0; i < 5 && index+i < p.TotalPage+1; i++ {
				p.Links = append(p.Links, index+i)
			}
		} else {
			for index := 1; index < p.TotalPage+1; index++ {
				p.Links = append(p.Links, index)
			}
		}

		p.Visable = true
	} else {
		p.Visable = false
	}
}

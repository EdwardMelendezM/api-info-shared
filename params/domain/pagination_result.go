package domain

type PaginationResults struct {
	CurrentPage int  `json:"current_page" binding:"required"`
	From        *int `json:"from"`
	LastPage    int  `json:"last_page" binding:"required"`
	SizePage    int  `json:"size_page" binding:"required"`
	To          *int `json:"to"`
	Total       int  `json:"total" binding:"required"`
}

func (p *PaginationResults) FromParams(pagination PaginationParams, total int) {
	p.Pagination(pagination.Page, pagination.SizePage, total)
}

func (p *PaginationResults) Pagination(page, sizePage, total int) {
	nPages := total / sizePage
	if total%sizePage > 0 {
		nPages++
	}
	from := (page-1)*sizePage + 1
	to := page * sizePage

	if to > total {
		to = total
	}

	if total == 0 {
		p.CurrentPage = page
		p.LastPage = 1
		p.SizePage = sizePage
		p.From = nil
		p.To = nil
		p.Total = total
	}
	p.CurrentPage = page
	p.LastPage = nPages
	p.SizePage = sizePage
	p.From = &from
	p.To = &to
	p.Total = total
}

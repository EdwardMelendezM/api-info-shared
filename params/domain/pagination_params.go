package domain

import (
	"net/http"
)

const (
	PageDefault     int = 1
	SizePageDefault int = 100
)

type PaginationParams struct {
	Params
	Page     int `json:"page"`
	SizePage int `json:"size_page"`
}

func NewPaginationParams(req *http.Request) PaginationParams {
	pagination := PaginationParams{}
	if req != nil {
		pagination.GetQueryParams(req, &pagination)
	}
	if pagination.Page == 0 {
		pagination.Page = PageDefault
	}
	if pagination.SizePage == 0 {
		pagination.SizePage = SizePageDefault
	}
	return pagination
}

func (p *PaginationParams) GetSizePage() int {
	if p.SizePage == 0 {
		p.SizePage = SizePageDefault
	}
	return p.SizePage
}

func (p *PaginationParams) GetOffset() int {
	if p.Page == 0 {
		p.Page = PageDefault
	}
	return (p.Page - 1) * p.SizePage
}

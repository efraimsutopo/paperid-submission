package structs

type (
	Pagination struct {
		Limit      int         `json:"limit"`
		Page       int         `json:"page"`
		Sort       string      `json:"sort"`
		TotalRows  int64       `json:"totalRows"`
		TotalPages int         `json:"totalPages"`
		Rows       interface{} `json:"rows"`
	}
)

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at DESC"
	}
	return p.Sort
}

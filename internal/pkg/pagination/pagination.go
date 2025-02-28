package pagination

import "math"

type PaginationResponse struct {
	Page       int   `json:"page" example:"1"`
	PerPage    int   `json:"per_page" example:"20"`
	PageCount  int   `json:"page_count" example:"10"`
	TotalCount int64 `json:"total_count" example:"200"`
}

func GeneratePaginationResponse(perPage, page int, total int64) (resp PaginationResponse) {
	pageCount := float64(total) / float64(perPage)
	resp = PaginationResponse{
		Page:       page,
		PerPage:    perPage,
		PageCount:  int(math.Ceil(pageCount)),
		TotalCount: total,
	}
	return
}

package response

import (
	"math"
)

type (
	MapResponse struct {
		Message  string      `json:"message"`
		Data     interface{} `json:"data"`
		Metadata *Pagination `json:"metadata,omitempty"`
	}

	Pagination struct {
		Page      int64 `json:"page"`
		Limit     int64 `json:"limit"`
		TotalData int   `json:"total_data"`
		TotalPage int   `json:"total_page"`
	}
)

const (
	SUCCESS = "success"
	CREATED = "created"
)

// CountTotalPage return total page of data
func (p *Pagination) CountTotalPage() {
	p.TotalPage = int(math.Ceil(float64(p.TotalData) / float64(p.Limit)))
}

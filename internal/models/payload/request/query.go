package request

import (
	"errors"
	"time"

	"github.com/baihakhi/to-do-list/internal/models/enum"
)

type (
	PaginationRequest struct {
		Page  int64  `json:"page" query:"page"`
		Limit int64  `json:"per_page" query:"per_page"`
		Key   string `json:"key" query:"q"`
	}
	DateRange struct {
		StartDate string `json:"start_date" query:"start_date"`
		EndDate   string `json:"end_date" query:"end_date"`
	}
	SearchBy enum.SearchBy
)

func (p *PaginationRequest) ValidatePagination() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = 10
	}
}

func (d *DateRange) ValidateDateRange() error {
	var (
		startDate, endDate time.Time
		err                error
	)

	if d.StartDate == "" && d.EndDate != "" {
		startDate = time.Now()
	} else if d.StartDate != "" && d.EndDate == "" {
		endDate = time.Now()
	} else if d.StartDate != "" && d.EndDate != "" {
		startDate, err = time.Parse(enum.DateFormatYYYYMMDD, d.StartDate)
		if err != nil {
			return err
		}
		endDate, err = time.Parse(enum.DateFormatYYYYMMDD, d.EndDate)
		if err != nil {
			return err
		}
	}

	if endDate.Before(startDate) {
		return errors.New("invalid date range")
	}

	return err
}

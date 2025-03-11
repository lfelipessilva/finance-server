package vo

import (
	"errors"
	"time"
)

type MonthYear struct {
	Month int
	Year  int
}

func NewMonthYear(month, year int) (MonthYear, error) {
	if month < 1 || month > 12 {
		return MonthYear{}, errors.New("invalid month")
	}
	if year < 2000 || year > 2100 {
		return MonthYear{}, errors.New("invalid year")
	}
	return MonthYear{Month: month, Year: year}, nil
}

func (my MonthYear) TimeRange() (start, end time.Time) {
	start = time.Date(my.Year, time.Month(my.Month), 1, 0, 0, 0, 0, time.UTC)
	end = start.AddDate(0, 1, 0)
	return start, end
}

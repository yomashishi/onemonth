package onemonth

import (
	"fmt"
	"time"
)

// During has BeginDay and EndDay.
type During struct {
	BeginDay time.Time
	EndDay   time.Time
}

// NewOneMonth returns a During struct with BeginDay and EndDay.
func NewOneMonth(year, month int) (*During, error) {
	if month < 1 || 12 < month {
		return nil, fmt.Errorf("month must be between 1 and 12, got: %d", month)
	}

	begin := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	nextMonth := int(begin.Month()) + 1
	end := time.Date(year, time.Month(nextMonth), 0, 0, 0, 0, 0, time.UTC)

	return &During{
		BeginDay: begin,
		EndDay:   end,
	}, nil
}

// Iterate through the days of the During.
func (d During) Iterate(f func(day time.Time)) {
	for day := d.BeginDay; !day.After(d.EndDay); day = day.AddDate(0, 0, 1) {
		f(day)
	}
}

// BeginMonth returns the time.Month of BeginDay.
func (d During) BeginMonth() time.Month {
	return d.BeginDay.Month()
}

// EndMonth returns the time.Month of EndDay.
func (d During) EndMonth() time.Month {
	return d.EndDay.Month()
}

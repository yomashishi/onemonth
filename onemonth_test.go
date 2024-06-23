package onemonth

import (
	"fmt"
	"testing"
	"time"
)

var year = 2022

func TestNew(t *testing.T) {
	testCases := []struct {
		month   int
		wantErr bool
	}{
		{month: 0, wantErr: true},
		{month: 1, wantErr: false},
		{month: 2, wantErr: false},
		{month: 13, wantErr: true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("month %d", tc.month), func(t *testing.T) {
			_, err := NewOneMonth(year, tc.month)
			if tc.wantErr {
				if err == nil {
					t.Error("no error occurred")
				}
			} else {
				if err != nil {
					t.Errorf("error occurred: %s", err)
				}
			}
		})
	}
}

func TestSameMonth(t *testing.T) {
	for i := 1; i <= 12; i++ {
		month, err := NewOneMonth(year, i)
		if err != nil {
			t.Errorf("failed to New, args: %d, %d", year, i)
		}

		if int(month.BeginMonth()) != i {
			t.Errorf("month.BeginMonth() and %d are not equal\n\tmonth.BeginMonth(): %d", i, month.BeginMonth())
		}
		if int(month.EndMonth()) != i {
			t.Errorf("month.EndMonth() and %d are not equal\n\tmonth.EndMonth(): %d", i, month.EndMonth())
		}
		if month.BeginMonth() != month.EndMonth() {
			t.Errorf("month.BeginMonth() and month.EndMonth() are not equal\n\tmonth.BeginMonth(): %d\n\tmonth.EndMonth(): %d",
				month.BeginMonth(), month.EndMonth(),
			)
		}
	}
}

func TestIterateCount(t *testing.T) {
	dayCountErr := "dayCount is not %d, year: %d, month: %d, dayCount: %d"
	days28months := map[int]struct{}{
		2: {},
	}
	days30months := map[int]struct{}{
		4:  {},
		6:  {},
		9:  {},
		11: {},
	}
	// 2024 is leap year.
	testYear := []int{2022, 2024}

	for _, year := range testYear {
		for i := 1; i <= 12; i++ {
			month, err := NewOneMonth(year, i)
			if err != nil {
				t.Errorf("failed to New, args: %d, %d", year, i)
			}

			dayCount := 0
			isLeapYear := false
			if time.Date(year, time.December, 31, 0, 0, 0, 0, time.Local).YearDay() > 365 {
				isLeapYear = true
			}

			month.Iterate(func(day time.Time) {
				dayCount++
			})

			if _, ok := days28months[i]; ok {
				if isLeapYear {
					if dayCount != 29 {
						t.Errorf(dayCountErr, 29, year, i, dayCount)
					}
				} else if dayCount != 28 {
					t.Errorf(dayCountErr, 28, year, i, dayCount)
				}
			} else if _, ok := days30months[i]; ok {
				if dayCount != 30 {
					t.Errorf(dayCountErr, 30, year, i, dayCount)
				}
			} else if dayCount != 31 {
				t.Errorf(dayCountErr, 31, year, i, dayCount)
			}
		}

	}
}

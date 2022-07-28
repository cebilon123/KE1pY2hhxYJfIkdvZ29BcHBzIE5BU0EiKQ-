package date_range

import "time"

// GetDateRangeSlice returns slice of dates between
// from and to date.
func GetDateRangeSlice(from, to time.Time) []time.Time {
	dates := []time.Time{}
	for dateIterate := dateRange(from, to); ; {
		date, isNext := dateIterate()

		if !isNext {
			break
		}

		dates = append(dates, date)
	}

	return dates
}

// dateRange iterates through
// range between date start and date end.
// returns func which returns two values:
// date and isNext. If isNext is false it
// means that there are no more dates
// to iterate on.
func dateRange(start, end time.Time) func() (time.Time, bool) {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() (time.Time, bool) {
		if start.After(end) || start.Equal(end) {
			return time.Time{}, false
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date, true
	}
}

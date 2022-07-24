package iterate

import "time"

func DateRange(start, end time.Time) func() (time.Time, bool) {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() (time.Time, bool) {
		if start.After(end) {
			return time.Time{}, false
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date, true
	}
}

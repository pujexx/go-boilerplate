package lib

import "time"

func DateFormat(dateString string) time.Time {
	t, e := time.Parse("2006-01-02", dateString)
	if e == nil {
		return t
	}
	return time.Now().UTC()
}

func DateRange(from time.Time, to time.Time) int {
	return int(to.Sub(from).Hours() / 24)
}

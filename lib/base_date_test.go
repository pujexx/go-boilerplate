package lib

import (
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	timeExpect := time.Date(2020, time.Month(5), 29, 0, 0, 0, 0, time.UTC)
	testResult := DateFormat("2020-05-29")
	assert.Equal(t, timeExpect.Month(), testResult.Month())
	assert.Equal(t, timeExpect.Year(), testResult.Year())
	assert.Equal(t, timeExpect.Day(), testResult.Day())

}

func TestDateRange(t *testing.T) {
	timeExpectFrom := time.Date(2020, time.Month(5), 29, 0, 0, 0, 0, time.UTC)
	timeExpectTo := time.Date(2020, time.Month(6), 1, 0, 0, 0, 0, time.UTC)
	day := DateRange(timeExpectFrom, timeExpectTo)
	assert.Equal(t, day, 3)
}

func TestDateRangeNegatife(t *testing.T) {
	timeExpectTo := time.Date(2020, time.Month(5), 29, 0, 0, 0, 0, time.UTC)
	timeExpectFrom := time.Date(2020, time.Month(6), 1, 0, 0, 0, 0, time.UTC)
	day := DateRange(timeExpectFrom, timeExpectTo)
	assert.Equal(t, day, -3)
}

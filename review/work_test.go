package review

import (
	"testing"
	"time"
)

const (
	year        = 2023
	month       = time.Month(8)
	day         = 1
	defaultTime = time.Duration(34200000000000)
)

func TestWork(t *testing.T) {

}

func getTodayDiff(startH, startM, endH, endM int) time.Duration {
	todayStart := time.Date(year, month, day, startH, startM, 0, 0, time.UTC)
	todayEnd := time.Date(year, month, day, endH, endM, 0, 0, time.UTC)
	return todayEnd.Sub(todayStart) - defaultTime
}

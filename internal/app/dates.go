package app

import (
	"fmt"
	"strconv"
	"time"

	"seman/internal/domain"
)

func weekStartOf(t time.Time) time.Time {
	year, month, day := t.Date()
	start := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	weekday := int(start.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return start.AddDate(0, 0, -(weekday - 1))
}

func weekLabel(start time.Time, span int) string {
	if span < 0 {
		return "All Weeks"
	}
	if span <= 1 {
		weekNum, _ := start.ISOWeek()
		end := start.AddDate(0, 0, 6)
		return "Week " + strconv.Itoa(weekNum) + " - " + domain.FormatDate(start) + " - " + domain.FormatDate(end)
	}
	end := start.AddDate(0, 0, span*7-1)
	startWeek, _ := start.ISOWeek()
	endWeek, _ := end.ISOWeek()
	return fmt.Sprintf("Weeks %d-%d - %s - %s", startWeek, endWeek, domain.FormatDate(start), domain.FormatDate(end))
}

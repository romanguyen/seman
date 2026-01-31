package domain

import (
	"strings"
	"time"
)

const DateLayout = "02/01/2006"
const DateTimeLayout = "02/01/2006 15:04"

func FormatDate(value time.Time) string {
	return value.Format(DateLayout)
}

func ParseStrictDate(value string) (time.Time, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}, false
	}
	if t, err := time.ParseInLocation(DateLayout, value, time.Local); err == nil {
		return t, true
	}
	return time.Time{}, false
}

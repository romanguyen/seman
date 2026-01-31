package domain

import (
	"strings"
	"time"
)

func ParseTodoDate(value string) (time.Time, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}, false
	}
	if t, err := time.ParseInLocation(DateLayout, value, time.Local); err == nil {
		return t, true
	}
	return time.Time{}, false
}

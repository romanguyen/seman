package app

import (
	"strings"

	"student-exams-manager/internal/domain"
)

func projectKey(items []domain.ProjectItem, idx int) string {
	if idx < 0 || idx >= len(items) {
		return ""
	}
	item := items[idx]
	return strings.ToUpper(item.Name) + "|" + strings.ToUpper(item.Subject) + "|" + item.Due
}

func findProjectIndex(items []domain.ProjectItem, key string) int {
	for i := range items {
		if projectKey(items, i) == key {
			return i
		}
	}
	return -1
}

package app

import (
	"sort"
	"strings"

	"student-exams-manager/internal/domain"
)

func (m *Model) sortProjectsByStatus() {
	if len(m.projects) < 2 {
		return
	}
	selectedKey := projectKey(m.projects, m.projectCursor)
	sort.SliceStable(m.projects, func(i, j int) bool {
		return statusRank(m.projects[i].Status) < statusRank(m.projects[j].Status)
	})
	if selectedKey != "" {
		if idx := findProjectIndex(m.projects, selectedKey); idx >= 0 {
			m.projectCursor = idx
		}
	}
}

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

func statusRank(status string) int {
	switch strings.ToUpper(strings.TrimSpace(status)) {
	case domain.ProjectStatusNotStarted:
		return 0
	case domain.ProjectStatusInProgress:
		return 1
	case domain.ProjectStatusDone:
		return 2
	default:
		return 3
	}
}

package app

import (
	"sort"
	"strings"

	"student-exams-manager/internal/domain"
)

func (m *Model) refreshExamFilter() {
	exams := m.examsForSelected()
	start, end, all := m.weekRange()
	visible := make([]int, 0, len(exams))
	if all {
		for i := range exams {
			visible = append(visible, i)
		}
	} else {
		for i, exam := range exams {
			date, ok := domain.ParseExamDate(exam.Date)
			if !ok {
				continue
			}
			if !date.Before(start) && date.Before(end) {
				visible = append(visible, i)
			}
		}
	}
	m.examVisible = visible
	m.normalizeExamCursor()
}

func (m *Model) filteredExams() []domain.ExamItem {
	exams := m.examsForSelected()
	if len(m.examVisible) == 0 {
		return nil
	}
	filtered := make([]domain.ExamItem, 0, len(m.examVisible))
	for _, idx := range m.examVisible {
		if idx >= 0 && idx < len(exams) {
			filtered = append(filtered, exams[idx])
		}
	}
	return filtered
}

func (m *Model) examsForSelected() []domain.ExamItem {
	if m.selectedSubj < 0 || m.selectedSubj >= len(m.subjects) {
		return nil
	}
	return m.subjects[m.selectedSubj].Exams
}

func (m *Model) normalizeExamCursor() {
	if len(m.examVisible) == 0 {
		m.examCursor = -1
		m.semesterFocus = focusSubjects
		return
	}
	if m.examCursor < 0 {
		m.examCursor = m.examVisible[0]
		return
	}
	if m.visibleIndex(m.examVisible, m.examCursor) == -1 {
		m.examCursor = m.examVisible[0]
	}
}

func (m *Model) sortExamsByPriority() {
	if len(m.subjects) == 0 {
		return
	}
	var selectedKey string
	selectedSubject := m.selectedSubj
	if selectedSubject >= 0 && selectedSubject < len(m.subjects) {
		selectedKey = examKey(m.subjects[selectedSubject].Exams, m.examCursor)
	}
	for i := range m.subjects {
		exams := m.subjects[i].Exams
		if len(exams) < 2 {
			continue
		}
		sort.SliceStable(exams, func(a, b int) bool {
			return priorityRank(exams[a].Priority) < priorityRank(exams[b].Priority)
		})
		m.subjects[i].Exams = exams
	}
	if selectedKey != "" && selectedSubject >= 0 && selectedSubject < len(m.subjects) {
		if idx := findExamIndex(m.subjects[selectedSubject].Exams, selectedKey); idx >= 0 {
			m.examCursor = idx
		}
	}
	m.refreshExamFilter()
}

func examKey(items []domain.ExamItem, idx int) string {
	if idx < 0 || idx >= len(items) {
		return ""
	}
	item := items[idx]
	return strings.ToUpper(item.Name) + "|" + item.Date + "|" + strings.ToUpper(item.Priority)
}

func findExamIndex(items []domain.ExamItem, key string) int {
	for i := range items {
		if examKey(items, i) == key {
			return i
		}
	}
	return -1
}

func priorityRank(priority string) int {
	switch strings.ToUpper(strings.TrimSpace(priority)) {
	case domain.PriorityHigh:
		return 0
	case domain.PriorityMed:
		return 1
	case domain.PriorityLow:
		return 2
	default:
		return 3
	}
}

func (m *Model) moveSemesterCursor(delta int) {
	if m.semesterFocus == focusSubjects {
		if len(m.subjects) == 0 {
			return
		}
		m.selectedSubj += delta
		if m.selectedSubj < 0 {
			m.selectedSubj = 0
		}
		if m.selectedSubj >= len(m.subjects) {
			m.selectedSubj = len(m.subjects) - 1
		}
		m.refreshExamFilter()
		return
	}

	if len(m.examVisible) == 0 {
		m.semesterFocus = focusSubjects
		return
	}
	pos := m.visibleIndex(m.examVisible, m.examCursor)
	if pos < 0 {
		pos = 0
	}
	pos += delta
	if pos < 0 {
		pos = 0
	}
	if pos >= len(m.examVisible) {
		pos = len(m.examVisible) - 1
	}
	m.examCursor = m.examVisible[pos]
}

func (m *Model) toggleSemesterFocus() {
	if m.semesterFocus == focusSubjects {
		if len(m.examVisible) == 0 {
			return
		}
		m.semesterFocus = focusExams
		return
	}
	m.semesterFocus = focusSubjects
}

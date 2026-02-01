package app

import "seman/internal/domain"

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
		if inBounds(idx, len(exams)) {
			filtered = append(filtered, exams[idx])
		}
	}
	return filtered
}

func (m *Model) examsForSelected() []domain.ExamItem {
	if !inBounds(m.selectedSubj, len(m.subjects)) {
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

func (m *Model) moveSemesterCursor(delta int) {
	if m.semesterFocus == focusSubjects {
		if len(m.subjects) == 0 {
			return
		}
		m.selectedSubj = clampIndex(m.selectedSubj+delta, len(m.subjects))
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
	pos = clampIndex(pos+delta, len(m.examVisible))
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

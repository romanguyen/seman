package app

import (
	"time"

	"seman/internal/domain"
)

func (m *Model) setWeekStartFromData(value string) {
	if value != "" {
		if t, ok := domain.ParseTodoDate(value); ok {
			m.weekStart = weekStartOf(t)
			m.updateWeekLabel()
			return
		}
	}
	m.weekStart = weekStartOf(time.Now())
	m.updateWeekLabel()
}

func (m *Model) shiftWeek(delta int) {
	m.weekStart = m.weekStart.AddDate(0, 0, delta*7)
	m.updateWeekLabel()
	m.refreshExamFilter()
	m.refreshChecklistView()
	m.persist()
}

func (m *Model) updateWeekLabel() {
	m.weekLabel = weekLabel(m.weekStart, m.weekSpan)
}

func (m *Model) setWeekSpanFromData(value int) {
	switch value {
	case -1:
		m.weekSpan = -1
	case 1, 2, 3, 4:
		m.weekSpan = value
	default:
		m.weekSpan = 1
	}
}

func (m *Model) cycleWeekSpan() {
	switch m.weekSpan {
	case 1:
		m.weekSpan = 2
	case 2:
		m.weekSpan = 3
	case 3:
		m.weekSpan = 4
	case 4:
		m.weekSpan = -1
	default:
		m.weekSpan = 1
	}
	m.updateWeekLabel()
	m.refreshExamFilter()
	m.refreshChecklistView()
	m.persist()
}

func (m *Model) weekRange() (time.Time, time.Time, bool) {
	if m.weekSpan < 0 {
		return time.Time{}, time.Time{}, true
	}
	start := m.weekStart
	end := start.AddDate(0, 0, m.weekSpan*7)
	return start, end, false
}

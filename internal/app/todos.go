package app

import (
	"strings"

	"student-exams-manager/internal/domain"
	"student-exams-manager/internal/style"
	"student-exams-manager/internal/ui/components"
)

func (m *Model) refreshChecklistView() {
	m.refreshTodoFilter()
	t := style.NewTheme()
	filtered := make([]domain.ChecklistItem, 0, len(m.todoVisible))
	for _, idx := range m.todoVisible {
		if idx >= 0 && idx < len(m.checklistItems) {
			filtered = append(filtered, m.checklistItems[idx])
		}
	}
	visibleCursor := m.visibleIndex(m.todoVisible, m.checklistCursor)
	m.checklist.SetContent(components.RenderChecklist(filtered, visibleCursor, m.activeTab == tabTodos, t))
	m.ensureChecklistVisible(visibleCursor, len(filtered))
}

func (m *Model) sortChecklistByDone() {
	if len(m.checklistItems) < 2 {
		return
	}

	oldCursor := m.checklistCursor
	if oldCursor < 0 {
		oldCursor = -1
	}
	indices := make([]int, 0, len(m.checklistItems))
	for i, item := range m.checklistItems {
		if !item.Done {
			indices = append(indices, i)
		}
	}
	for i, item := range m.checklistItems {
		if item.Done {
			indices = append(indices, i)
		}
	}

	newItems := make([]domain.ChecklistItem, len(m.checklistItems))
	newCursor := -1
	for newIdx, oldIdx := range indices {
		newItems[newIdx] = m.checklistItems[oldIdx]
		if oldIdx == oldCursor {
			newCursor = newIdx
		}
	}

	m.checklistItems = newItems
	m.checklistCursor = newCursor
}

func (m *Model) ensureChecklistVisible(cursor, count int) {
	if count == 0 || m.checklist.Height <= 0 {
		return
	}
	if cursor < m.checklist.YOffset {
		m.checklist.SetYOffset(cursor)
		return
	}
	if cursor >= m.checklist.YOffset+m.checklist.Height {
		m.checklist.SetYOffset(cursor - m.checklist.Height + 1)
	}
}

func (m *Model) moveChecklistCursor(delta int) {
	if len(m.todoVisible) == 0 {
		return
	}
	pos := m.visibleIndex(m.todoVisible, m.checklistCursor)
	if pos < 0 {
		pos = 0
	}
	pos += delta
	if pos < 0 {
		pos = 0
	}
	if pos >= len(m.todoVisible) {
		pos = len(m.todoVisible) - 1
	}
	m.checklistCursor = m.todoVisible[pos]
	m.refreshChecklistView()
}

func (m *Model) toggleChecklistItem() {
	if len(m.todoVisible) == 0 {
		return
	}
	idx := m.checklistCursor
	if idx < 0 || idx >= len(m.checklistItems) {
		return
	}
	m.checklistItems[idx].Done = !m.checklistItems[idx].Done
	m.sortChecklistByDone()
	m.persist()
	m.refreshChecklistView()
}

func (m *Model) refreshTodoFilter() {
	start, end, all := m.weekRange()
	visible := make([]int, 0, len(m.checklistItems))
	if all {
		for i := range m.checklistItems {
			visible = append(visible, i)
		}
	} else {
		for i, item := range m.checklistItems {
			due, ok := domain.ParseTodoDate(item.Due)
			if !ok {
				continue
			}
			if !due.Before(start) && due.Before(end) {
				visible = append(visible, i)
			}
		}
	}
	m.todoVisible = visible
	if len(m.todoVisible) == 0 {
		m.checklistCursor = -1
		return
	}
	if m.visibleIndex(m.todoVisible, m.checklistCursor) == -1 {
		m.checklistCursor = m.todoVisible[0]
	}
}

func (m *Model) ensureTodoDueDates() {
	if len(m.checklistItems) == 0 {
		return
	}
	changed := false
	defaultDue := m.weekStart.Format("2006-01-02")
	for i := range m.checklistItems {
		if strings.TrimSpace(m.checklistItems[i].Due) == "" {
			m.checklistItems[i].Due = defaultDue
			changed = true
		}
	}
	if changed {
		m.persist()
	}
}

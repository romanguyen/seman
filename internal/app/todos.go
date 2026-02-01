package app

import (
	"strings"

	"seman/internal/domain"
	"seman/internal/style"
	"seman/internal/ui/components"
)

func (m *Model) refreshChecklistView() {
	m.refreshTodoFilter()
	t := style.NewTheme()
	filtered := make([]domain.ChecklistItem, 0, len(m.todoVisible))
	for _, idx := range m.todoVisible {
		if inBounds(idx, len(m.checklistItems)) {
			filtered = append(filtered, m.checklistItems[idx])
		}
	}
	visibleCursor := m.visibleIndex(m.todoVisible, m.checklistCursor)
	m.checklist.SetContent(components.RenderChecklist(filtered, visibleCursor, m.activeTab == tabTodos, t))
	m.ensureChecklistVisible(visibleCursor, len(filtered))
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
	pos = clampIndex(pos+delta, len(m.todoVisible))
	m.checklistCursor = m.todoVisible[pos]
	m.refreshChecklistView()
}

func (m *Model) toggleChecklistItem() {
	if len(m.todoVisible) == 0 {
		return
	}
	idx := m.checklistCursor
	if !inBounds(idx, len(m.checklistItems)) {
		return
	}
	m.checklistItems[idx].Done = !m.checklistItems[idx].Done
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
	defaultDue := domain.FormatDate(m.weekStart)
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

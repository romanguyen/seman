package app

import "student-exams-manager/internal/ui/layout"

func (m *Model) resize(width, height int) {
	m.width = width
	m.height = height

	mainHeight := layout.MainAreaHeight(height)
	switch m.activeTab {
	case tabDashboard:
		bounds := layout.ComputeDashboardLayout(width, mainHeight)
		if bounds.ChecklistWidth > 0 && bounds.ChecklistHeight > 0 {
			m.checklist.Width = bounds.ChecklistWidth
			m.checklist.Height = bounds.ChecklistHeight
		}
	case tabTodos:
		bounds := layout.ComputeWeeklyLayout(width, mainHeight)
		if bounds.ActionsWidth > 0 && bounds.ActionsHeight > 0 {
			m.checklist.Width = bounds.ActionsWidth
			m.checklist.Height = bounds.ActionsHeight
		}
	case tabLofi:
		m.resizeLofi(width, mainHeight)
	}
	m.refreshChecklistView()
}

func (m *Model) resizeLofi(width, height int) {
	gap := 1
	available := width - gap
	if available < 1 {
		available = width
		gap = 0
	}
	leftWidth := available / 2
	rightWidth := available - leftWidth
	if leftWidth < 1 {
		leftWidth = available
		rightWidth = 0
	}
	playlistWidth := rightWidth
	if rightWidth == 0 {
		m.lofiListHeight = 0
		m.lofiOffset = 0
		return
	}

	_, contentH := layout.PanelContentSize(playlistWidth, height)
	if contentH < 1 {
		contentH = 1
	}
	m.lofiListHeight = contentH
	m.ensureLofiVisible()
}

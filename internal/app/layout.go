package app

import "seman/internal/ui/layout"

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
		bounds := layout.ComputeTodoLayout(width, mainHeight)
		if bounds.ActionsWidth > 0 && bounds.ActionsHeight > 0 {
			m.checklist.Width = bounds.ActionsWidth
			m.checklist.Height = bounds.ActionsHeight
		}
	case tabLofi:
		return
	}
	m.refreshChecklistView()
}

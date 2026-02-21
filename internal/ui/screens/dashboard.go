package screens

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/style"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/ui/components"
)

func RenderDashboard(state State, width, height int, t style.Theme) string {
	layout := components.ComputeDashboardLayout(width, height)

	leftBody := components.RenderUpcomingExams(state.Subjects, 5, state.FilterStart, state.FilterEnd, state.FilterAll, t)
	leftPanel := components.RenderPanel(layout.LeftWidth, layout.PanelHeight, "Upcoming Exams (Priority)", leftBody, t)

	if layout.MiddleWidth <= 0 {
		return leftPanel
	}

	middlePanel := components.RenderPanel(layout.MiddleWidth, layout.PanelHeight, "Todos", state.ChecklistView, t)
	if layout.RightWidth <= 0 {
		return lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, " ", middlePanel)
	}

	rightPanel := components.RenderPanel(layout.RightWidth, layout.PanelHeight, "Projects (Soonest Deadlines)", components.RenderProjects(state.Projects, t), t)
	return lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, " ", middlePanel, " ", rightPanel)
}

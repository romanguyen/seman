package dashboard

import (
	"github.com/charmbracelet/lipgloss"
	"student-exams-manager/internal/screens"
	"student-exams-manager/internal/style"
	"student-exams-manager/internal/ui/components"
	"student-exams-manager/internal/ui/layout"
)

func Render(state screens.State, width, height int, t style.Theme) string {
	bounds := layout.ComputeDashboardLayout(width, height)

	leftBody := components.RenderUpcomingExams(state.Subjects, 5, state.FilterStart, state.FilterEnd, state.FilterAll, t)
	leftPanel := components.RenderPanel(bounds.LeftWidth, bounds.PanelHeight, "Upcoming Exams (Priority)", leftBody, t)

	if bounds.MiddleWidth <= 0 {
		return leftPanel
	}

	middlePanel := components.RenderPanel(bounds.MiddleWidth, bounds.PanelHeight, "Todos", state.Checklist.View(), t)
	if bounds.RightWidth <= 0 {
		return lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, " ", middlePanel)
	}

	rightPanel := components.RenderPanel(bounds.RightWidth, bounds.PanelHeight, "Projects (Soonest Deadlines)", components.RenderProjects(state.Projects, t), t)
	return lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, " ", middlePanel, " ", rightPanel)
}

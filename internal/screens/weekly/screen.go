package weekly

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"student-exams-manager/internal/screens"
	"student-exams-manager/internal/style"
	"student-exams-manager/internal/ui/components"
	"student-exams-manager/internal/ui/layout"
)

func Render(state screens.State, width, height int, t style.Theme) string {
	bounds := layout.ComputeWeeklyLayout(width, height)

	var parts []string
	rightPanel := components.RenderPanel(bounds.RightWidth, bounds.PanelsHeight, "Todos", state.Checklist.View(), t)
	parts = append(parts, rightPanel)

	if bounds.SpacerHeight > 0 {
		parts = append(parts, strings.Repeat("\n", bounds.SpacerHeight-1))
	}

	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

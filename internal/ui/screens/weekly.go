package screens

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/romanguyen/seman/internal/style"
	"github.com/romanguyen/seman/internal/ui/components"
)

func RenderWeeklyFocus(state State, width, height int, t style.Theme) string {
	layout := components.ComputeWeeklyLayout(width, height)

	var parts []string
	rightPanel := components.RenderPanel(layout.RightWidth, layout.PanelsHeight, "Todos", state.ChecklistView, t)
	parts = append(parts, rightPanel)

	if layout.SpacerHeight > 0 {
		parts = append(parts, strings.Repeat("\n", layout.SpacerHeight-1))
	}

	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

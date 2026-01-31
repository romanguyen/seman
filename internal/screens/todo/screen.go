package todo

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"seman/internal/screens"
	"seman/internal/style"
	"seman/internal/ui/components"
	"seman/internal/ui/layout"
)

func Render(state screens.State, width, height int, t style.Theme) string {
	bounds := layout.ComputeTodoLayout(width, height)

	var parts []string
	rightPanel := components.RenderPanel(bounds.RightWidth, bounds.PanelsHeight, "Todos", state.Checklist.View(), t)
	parts = append(parts, rightPanel)

	if bounds.SpacerHeight > 0 {
		parts = append(parts, strings.Repeat("\n", bounds.SpacerHeight-1))
	}

	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

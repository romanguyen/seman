package screens

import (
	"github.com/charmbracelet/lipgloss"
	"seman/internal/style"
	"seman/internal/ui/components"
)

func RenderPlaceholder(width, height int, t style.Theme) string {
	message := t.Dim.Render("Coming soon")
	box := lipgloss.NewStyle().
		Width(width).
		Height(height).
		Align(lipgloss.Center, lipgloss.Center)
	return box.Render(message)
}

func RenderModal(state State, width, height int, t style.Theme) string {
	return components.RenderModalArea(state.Modal, width, height, t)
}

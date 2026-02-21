package screens

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/style"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/ui/components"
)

func RenderMain(state State, width, height int, t style.Theme) string {
	switch state.ActiveTab {
	case 0:
		return RenderDashboard(state, width, height, t)
	case 1:
		return RenderSemester(state, width, height, t)
	case 2:
		return RenderWeeklyFocus(state, width, height, t)
	case 3:
		return RenderProjectsTab(state, width, height, t)
	case 4:
		return RenderSettingsTab(width, height, state.ConfirmOn, state.WeekSpan, state.LofiEnabled, state.LofiURL, t)
	case 5:
		return RenderLofi(state, width, height, t)
	default:
		return RenderPlaceholder(width, height, t)
	}
}

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

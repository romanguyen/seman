package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"seman/internal/domain"
	"seman/internal/style"
)

func RenderStatusBadge(status string, width int, t style.Theme) string {
	label := strings.ToUpper(status)
	style := t.StatusNotStr
	switch label {
	case domain.ProjectStatusDone:
		style = t.StatusDone
	case domain.ProjectStatusInProgress:
		style = t.StatusInProg
	case domain.ProjectStatusNotStarted:
		style = t.StatusNotStr
	}
	badge := style.Render(label)
	if lipgloss.Width(badge) < width {
		return badge + strings.Repeat(" ", width-lipgloss.Width(badge))
	}
	return TruncateString(badge, width)
}

func RenderPriority(priority string) string {
	style := lipgloss.NewStyle().Padding(0, 1).Bold(true)
	switch strings.ToUpper(priority) {
	case domain.PriorityHigh:
		style = style.Foreground(lipgloss.Color("#0c0c0c")).Background(lipgloss.Color("#ff5f5f"))
	case domain.PriorityMed:
		style = style.Foreground(lipgloss.Color("#0c0c0c")).Background(lipgloss.Color("#d4a017"))
	case domain.PriorityLow:
		style = style.Foreground(lipgloss.Color("#0c0c0c")).Background(lipgloss.Color("#3a7f3a"))
	default:
		style = style.Foreground(lipgloss.Color("#0c0c0c")).Background(lipgloss.Color("#39ff14"))
	}
	return style.Render(priority)
}

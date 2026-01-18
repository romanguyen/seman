package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"student-exams-manager/internal/style"
)

func RenderTabs(active, width int, weekLabel string, t style.Theme) string {
	labels := []string{"Dashboard", "Exams", "Todos", "Projects", "Settings"}
	tabs := make([]string, 0, len(labels))
	for i, label := range labels {
		text := fmt.Sprintf("[%d] %s", i+1, label)
		if i == active {
			tabs = append(tabs, t.TabActive.Render(text))
		} else {
			tabs = append(tabs, t.TabInactive.Render(text))
		}
	}
	left := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
	left = lipgloss.NewStyle().PaddingLeft(1).Render(left)

	if weekLabel == "" || width <= lipgloss.Width(left) {
		return left
	}

	rightLabel := "< " + weekLabel + " >"
	right := t.Dim.Render(rightLabel)
	rightWidth := width - lipgloss.Width(left)
	if rightWidth < 1 {
		return left
	}
	right = lipgloss.NewStyle().Width(rightWidth).Height(3).Align(lipgloss.Right, lipgloss.Center).Render(right)
	if lipgloss.Width(right) > rightWidth {
		right = strings.TrimSpace(lipgloss.NewStyle().Width(rightWidth).Render(right))
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

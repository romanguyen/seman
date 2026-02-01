package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func AlignLine(width int, left, right string) string {
	leftWidth := lipgloss.Width(left)
	rightWidth := lipgloss.Width(right)
	gap := width - leftWidth - rightWidth
	if gap < 1 {
		gap = 1
	}
	return left + strings.Repeat(" ", gap) + right
}

func TruncateString(s string, width int) string {
	if width <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= width {
		return s
	}
	return string(runes[:width])
}

func YesNo(val bool) string {
	if val {
		return "Yes"
	}
	return "No"
}

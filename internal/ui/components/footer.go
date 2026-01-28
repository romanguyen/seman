package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"student-exams-manager/internal/style"
	"student-exams-manager/internal/ui/layout"
)

func RenderFooter(width, tabCount int, status string, t style.Theme) string {
	contentWidth := width - layout.BarBorderX - layout.BarPaddingX*2
	if contentWidth < 1 {
		contentWidth = 1
	}

	left := t.FooterHint.Render("[A] Add exam  [S] Add subject  [P] Add project  [Q] Quit")
	if status != "" {
		left = t.ModalError.Render(TruncateString(status, contentWidth))
	}
	right := t.FooterHint.Render(fmt.Sprintf("[1-%d] Switch tabs", tabCount))
	content := AlignLine(contentWidth, left, right)

	styleWidth := width - layout.BarBorderX
	if styleWidth < 1 {
		styleWidth = 1
	}

	box := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(t.Border).
		Padding(layout.BarPaddingY, layout.BarPaddingX).
		Width(styleWidth)

	return box.Render(content)
}

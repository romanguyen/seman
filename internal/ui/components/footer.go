package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"seman/internal/style"
	"seman/internal/ui/layout"
)

func RenderFooter(width, tabCount, activeTab int, status string, t style.Theme) string {
	contentWidth := width - layout.BarBorderX - layout.BarPaddingX*2
	if contentWidth < 1 {
		contentWidth = 1
	}

	hint := "[A] Add  [S] Add subject  [E] Edit  [D] Delete  [Q] Quit"
	if activeTab == 1 {
		hint += "  [Tab] Focus"
	}
	left := t.FooterHint.Render(hint)
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

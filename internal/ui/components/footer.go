package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/style"
)

func RenderFooter(width, tabCount int, t style.Theme) string {
	contentWidth := width - barBorderX - barPaddingX*2
	if contentWidth < 1 {
		contentWidth = 1
	}

	left := t.FooterHint.Render("[A] Add exam  [S] Add subject  [P] Add project  [Q] Quit")
	right := t.FooterHint.Render(fmt.Sprintf("[1-%d] Switch tabs", tabCount))
	content := AlignLine(contentWidth, left, right)

	styleWidth := width - barBorderX
	if styleWidth < 1 {
		styleWidth = 1
	}

	box := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(t.Border).
		Padding(barPaddingY, barPaddingX).
		Width(styleWidth)

	return box.Render(content)
}

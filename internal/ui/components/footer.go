package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/style"
)

// tabSubjects is the tab ID for the Subjects tab, kept in sync with app/model.go.
const tabSubjectsID = 6

func RenderFooter(width, tabCount, activeTab int, t style.Theme) string {
	contentWidth := width - barBorderX - barPaddingX*2
	if contentWidth < 1 {
		contentWidth = 1
	}

	var leftStr string
	if activeTab == tabSubjectsID {
		leftStr = "[A] Add subject  [E] Edit  [D] Delete  [Q] Quit"
	} else {
		leftStr = "[A] Add exam  [S] Add subject  [P] Add project  [Q] Quit"
	}
	left := t.FooterHint.Render(leftStr)
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

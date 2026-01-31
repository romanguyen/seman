package components

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	"student-exams-manager/internal/domain"
	"student-exams-manager/internal/style"
	"student-exams-manager/internal/ui/layout"
)

func RenderHeader(width int, t style.Theme) string {
	contentWidth := width - layout.BarBorderX - layout.BarPaddingX*2
	if contentWidth < 1 {
		contentWidth = 1
	}

	title := t.Title.Render("Student Manager")
	date := t.Dim.Render(domain.FormatDate(time.Now()))
	content := AlignLine(contentWidth, title, date)

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

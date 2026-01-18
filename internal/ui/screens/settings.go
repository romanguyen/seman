package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"student-exams-manager/internal/style"
	"student-exams-manager/internal/ui/components"
)

func RenderSettingsTab(width, height int, confirmOn bool, weekSpan int, t style.Theme) string {
	gap := 1
	colWidth := (width - gap) / 2
	rightWidth := width - colWidth - gap
	if colWidth < 1 {
		colWidth = width
		rightWidth = 0
	}

	dataBody := strings.Join([]string{
		t.Text.Render("[X] Export data to JSON"),
		t.Text.Render("[I] Import data from file"),
		t.Text.Render("[B] Backup current semester"),
		t.Text.Render("[C] Clear all data (CAUTION)"),
	}, "\n")

	displayContentW := components.PanelContentWidth(colWidth)
	if rightWidth > 0 {
		displayContentW = components.PanelContentWidth(rightWidth)
	}
	displayBody := strings.Join([]string{
		components.AlignLine(displayContentW, t.Text.Render("Weeks visible: "+weekSpanLabel(weekSpan)), t.Text.Render("[W] Change")),
		components.AlignLine(displayContentW, t.Text.Render("Date format: DD/MM/YYYY"), t.Text.Render("[F] Change")),
		components.AlignLine(displayContentW, t.Text.Render("Time format: 24-hour"), t.Text.Render("[T] Toggle")),
		components.AlignLine(displayContentW, t.Text.Render("Highlight urgent items: Yes"), t.Text.Render("[H] Toggle")),
		components.AlignLine(displayContentW, t.Text.Render(fmt.Sprintf("Confirm deletions: %s", components.YesNo(confirmOn))), t.Text.Render("[O] Toggle")),
	}, "\n")

	dataLines := strings.Count(dataBody, "\n") + 1 + 1
	displayLines := strings.Count(displayBody, "\n") + 1 + 1
	dataHeight := components.PanelHeightForLines(dataLines)
	displayHeight := components.PanelHeightForLines(displayLines)
	topRowHeight := dataHeight
	if displayHeight > topRowHeight {
		topRowHeight = displayHeight
	}

	priorityBody := strings.Join([]string{
		t.Text.Render("HIGH priority if exam is within: 7 days"),
		t.Text.Render("MEDIUM priority if exam is within: 14 days"),
		t.Text.Render("LOW priority otherwise"),
		t.Dim.Render("[R] Configure rules"),
	}, "\n")

	priorityLines := strings.Count(priorityBody, "\n") + 1 + 1
	priorityHeight := components.PanelHeightForLines(priorityLines)

	totalHeight := topRowHeight + priorityHeight + gap
	if extra := height - totalHeight; extra > 0 {
		priorityHeight += extra
	}

	dataPanel := components.RenderPanel(colWidth, topRowHeight, "Data Management", dataBody, t)
	var topRow string
	if rightWidth > 0 {
		displayPanel := components.RenderPanel(rightWidth, topRowHeight, "Display Options", displayBody, t)
		topRow = lipgloss.JoinHorizontal(lipgloss.Top, dataPanel, " ", displayPanel)
	} else {
		topRow = dataPanel
	}

	priorityPanel := components.RenderPanel(width, priorityHeight, "Priority Rules", priorityBody, t)

	return lipgloss.JoinVertical(lipgloss.Left, topRow, "", priorityPanel)
}

func weekSpanLabel(span int) string {
	switch span {
	case -1:
		return "All"
	case 1:
		return "1 week"
	case 2, 3, 4:
		return fmt.Sprintf("%d weeks", span)
	default:
		return "1 week"
	}
}

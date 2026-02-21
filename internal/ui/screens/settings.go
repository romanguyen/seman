package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/style"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/ui/components"
)

func RenderSettingsTab(width, height int, confirmOn bool, weekSpan int, lofiEnabled bool, lofiURL string, t style.Theme) string {
	gap := 1
	leftWidth := (width - gap) / 2
	rightWidth := width - leftWidth - gap
	if leftWidth < 1 {
		leftWidth = width
		rightWidth = 0
	}

	dataBody := strings.Join([]string{
		t.Text.Render("[X] Export data to JSON"),
		t.Text.Render("[I] Import data from file"),
		t.Text.Render("[B] Backup current semester"),
		t.Text.Render("[C] Clear all data (CAUTION)"),
	}, "\n")

	displayContentW := components.PanelContentWidth(leftWidth)
	displayBody := strings.Join([]string{
		components.AlignLine(displayContentW, t.Text.Render("Weeks visible: "+weekSpanLabel(weekSpan)), t.Text.Render("[W] Change")),
		components.AlignLine(displayContentW, t.Text.Render("Date format: DD/MM/YYYY"), t.Text.Render("[F] Change")),
		components.AlignLine(displayContentW, t.Text.Render("Time format: 24-hour"), t.Text.Render("[T] Toggle")),
		components.AlignLine(displayContentW, t.Text.Render("Highlight urgent items: Yes"), t.Text.Render("[H] Toggle")),
		components.AlignLine(displayContentW, t.Text.Render(fmt.Sprintf("Confirm deletions: %s", components.YesNo(confirmOn))), t.Text.Render("[O] Toggle")),
		components.AlignLine(displayContentW, t.Text.Render(fmt.Sprintf("Lofi tab: %s", components.YesNo(lofiEnabled))), t.Text.Render("[L] Toggle")),
		components.AlignLine(displayContentW, t.Text.Render("Lofi playlist: "+lofiURLLabel(lofiURL, displayContentW)), t.Text.Render("[U] Edit")),
	}, "\n")

	dataLines := strings.Count(dataBody, "\n") + 1 + 1
	dataHeight := components.PanelHeightForLines(dataLines)

	priorityBody := strings.Join([]string{
		t.Text.Render("HIGH priority if exam is within: 7 days"),
		t.Text.Render("MEDIUM priority if exam is within: 14 days"),
		t.Text.Render("LOW priority otherwise"),
		t.Dim.Render("[R] Configure rules"),
	}, "\n")

	priorityLines := strings.Count(priorityBody, "\n") + 1 + 1
	priorityHeight := components.PanelHeightForLines(priorityLines)

	totalRightHeight := dataHeight + priorityHeight + gap
	if extra := height - totalRightHeight; extra > 0 {
		priorityHeight += extra
	} else if extra < 0 {
		priorityHeight += extra
		if priorityHeight < 1 {
			priorityHeight = 1
		}
	}

	displayPanel := components.RenderPanel(leftWidth, height, "Display Options", displayBody, t)
	if rightWidth == 0 {
		return displayPanel
	}
	dataPanel := components.RenderPanel(rightWidth, dataHeight, "Data Management", dataBody, t)
	priorityPanel := components.RenderPanel(rightWidth, priorityHeight, "Priority Rules", priorityBody, t)
	rightColumn := lipgloss.JoinVertical(lipgloss.Left, dataPanel, "", priorityPanel)

	return lipgloss.JoinHorizontal(lipgloss.Top, displayPanel, " ", rightColumn)
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

func lofiURLLabel(url string, width int) string {
	label := strings.TrimSpace(url)
	if label == "" {
		label = "(not set)"
	}
	max := width - len("Lofi playlist: ")
	if max < 1 {
		max = 1
	}
	return components.TruncateString(label, max)
}

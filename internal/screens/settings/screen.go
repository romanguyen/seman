package settings

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"seman/internal/screens"
	"seman/internal/style"
	"seman/internal/ui/components"
	"seman/internal/ui/layout"
)

func themeLabel(name style.ThemeName) string {
	return style.ThemeLabel(name)
}

func Render(state screens.State, width, height int, t style.Theme) string {
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

	displayContentW := layout.PanelContentWidth(leftWidth)
	displayBody := strings.Join([]string{
		components.AlignLine(displayContentW, t.Text.Render("Color theme: "+themeLabel(state.ThemeName)), t.Text.Render("[T] Change")),
		components.AlignLine(displayContentW, t.Text.Render("Weeks visible: "+weekSpanLabel(state.WeekSpan)), t.Text.Render("[W] Change")),
		components.AlignLine(displayContentW, t.Text.Render(fmt.Sprintf("Confirm deletions: %s", components.YesNo(state.ConfirmOn))), t.Text.Render("[O] Toggle")),
		components.AlignLine(displayContentW, t.Text.Render(fmt.Sprintf("Lofi tab: %s", components.YesNo(state.LofiEnabled))), t.Text.Render("[L] Toggle")),
		components.AlignLine(displayContentW, t.Text.Render("Lofi playlist: "+lofiURLLabel(state.LofiURL, displayContentW)), t.Text.Render("[U] Edit")),
	}, "\n")

	displayPanel := components.RenderPanel(leftWidth, height, "Display Options", displayBody, t)
	if rightWidth == 0 {
		return displayPanel
	}
	dataPanel := components.RenderPanel(rightWidth, height, "Data Management", dataBody, t)

	return lipgloss.JoinHorizontal(lipgloss.Top, displayPanel, " ", dataPanel)
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

package style

import "github.com/charmbracelet/lipgloss"

// ThemeName represents a color theme identifier.
type ThemeName string

const (
	ThemeMatrix   ThemeName = "matrix"
	ThemeDracula  ThemeName = "dracula"
	ThemeNord     ThemeName = "nord"
	ThemeSolarize ThemeName = "solarized"
	ThemeCyberpnk ThemeName = "cyberpunk"
)

// ThemeNames returns all available theme names in order.
func ThemeNames() []ThemeName {
	return []ThemeName{ThemeMatrix, ThemeDracula, ThemeNord, ThemeSolarize, ThemeCyberpnk}
}

// ThemeLabel returns a display label for a theme name.
func ThemeLabel(name ThemeName) string {
	switch name {
	case ThemeMatrix:
		return "Matrix"
	case ThemeDracula:
		return "Dracula"
	case ThemeNord:
		return "Nord"
	case ThemeSolarize:
		return "Solarized"
	case ThemeCyberpnk:
		return "Cyberpunk"
	default:
		return "Matrix"
	}
}

// NextTheme returns the next theme in the cycle.
func NextTheme(current ThemeName) ThemeName {
	themes := ThemeNames()
	for i, t := range themes {
		if t == current {
			return themes[(i+1)%len(themes)]
		}
	}
	return ThemeMatrix
}

type themeColors struct {
	accent     lipgloss.Color
	border     lipgloss.Color
	rowBg      lipgloss.Color
	statusDone lipgloss.Color
	statusProg lipgloss.Color
	statusNot  lipgloss.Color
	statusFg   lipgloss.Color
}

func colorsForTheme(name ThemeName) themeColors {
	switch name {
	case ThemeDracula:
		return themeColors{
			accent:     lipgloss.Color("#bd93f9"),
			border:     lipgloss.Color("#6272a4"),
			rowBg:      lipgloss.Color("#282a36"),
			statusDone: lipgloss.Color("#50fa7b"),
			statusProg: lipgloss.Color("#ffb86c"),
			statusNot:  lipgloss.Color("#6272a4"),
			statusFg:   lipgloss.Color("#282a36"),
		}
	case ThemeNord:
		return themeColors{
			accent:     lipgloss.Color("#88c0d0"),
			border:     lipgloss.Color("#4c566a"),
			rowBg:      lipgloss.Color("#3b4252"),
			statusDone: lipgloss.Color("#a3be8c"),
			statusProg: lipgloss.Color("#ebcb8b"),
			statusNot:  lipgloss.Color("#4c566a"),
			statusFg:   lipgloss.Color("#2e3440"),
		}
	case ThemeSolarize:
		return themeColors{
			accent:     lipgloss.Color("#268bd2"),
			border:     lipgloss.Color("#586e75"),
			rowBg:      lipgloss.Color("#073642"),
			statusDone: lipgloss.Color("#859900"),
			statusProg: lipgloss.Color("#b58900"),
			statusNot:  lipgloss.Color("#586e75"),
			statusFg:   lipgloss.Color("#002b36"),
		}
	case ThemeCyberpnk:
		return themeColors{
			accent:     lipgloss.Color("#ff00ff"),
			border:     lipgloss.Color("#00ffff"),
			rowBg:      lipgloss.Color("#1a0a2e"),
			statusDone: lipgloss.Color("#00ff00"),
			statusProg: lipgloss.Color("#ffff00"),
			statusNot:  lipgloss.Color("#ff00ff"),
			statusFg:   lipgloss.Color("#0a0a0a"),
		}
	default: // ThemeMatrix
		return themeColors{
			accent:     lipgloss.Color("#39ff14"),
			border:     lipgloss.Color("#1e5f1e"),
			rowBg:      lipgloss.Color("#0b140b"),
			statusDone: lipgloss.Color("#3a7f3a"),
			statusProg: lipgloss.Color("#d4a017"),
			statusNot:  lipgloss.Color("#1e5f1e"),
			statusFg:   lipgloss.Color("#0c0c0c"),
		}
	}
}

type Theme struct {
	Name          ThemeName
	Accent        lipgloss.Color
	Border        lipgloss.Color
	Text          lipgloss.Style
	Dim           lipgloss.Style
	Title         lipgloss.Style
	TabActive     lipgloss.Style
	TabInactive   lipgloss.Style
	CheckboxDone  lipgloss.Style
	CheckboxTodo  lipgloss.Style
	FooterHint    lipgloss.Style
	ProjectDetail lipgloss.Style
	SubjectActive lipgloss.Style
	SubjectDim    lipgloss.Style
	StatusDone    lipgloss.Style
	StatusInProg  lipgloss.Style
	StatusNotStr  lipgloss.Style
	RowActive     lipgloss.Style
	ModalBorder   lipgloss.Style
	ModalTitle    lipgloss.Style
	ModalHint     lipgloss.Style
	ModalError    lipgloss.Style
	InputText     lipgloss.Style
	InputHint     lipgloss.Style
	InputCursor   lipgloss.Style
}

func NewTheme() Theme {
	return NewThemeWithName(ThemeMatrix)
}

func NewThemeWithName(name ThemeName) Theme {
	c := colorsForTheme(name)
	text := lipgloss.NewStyle().Foreground(c.accent)
	dim := lipgloss.NewStyle().Foreground(c.border)
	return Theme{
		Name:          name,
		Accent:        c.accent,
		Border:        c.border,
		Text:          text,
		Dim:           dim,
		Title:         text.Copy().Bold(true),
		TabActive:     lipgloss.NewStyle().Foreground(c.accent).Border(lipgloss.NormalBorder()).BorderForeground(c.accent).Padding(0, 1),
		TabInactive:   lipgloss.NewStyle().Foreground(c.border).Padding(1, 1),
		CheckboxDone:  dim.Copy(),
		CheckboxTodo:  text.Copy(),
		FooterHint:    dim.Copy(),
		ProjectDetail: dim.Copy(),
		SubjectActive: text.Copy().Bold(true),
		SubjectDim:    dim.Copy(),
		StatusDone:    lipgloss.NewStyle().Foreground(c.statusFg).Background(c.statusDone).Padding(0, 1),
		StatusInProg:  lipgloss.NewStyle().Foreground(c.statusFg).Background(c.statusProg).Padding(0, 1),
		StatusNotStr:  lipgloss.NewStyle().Foreground(c.statusFg).Background(c.statusNot).Padding(0, 1),
		RowActive:     lipgloss.NewStyle().Foreground(c.accent).Background(c.rowBg).Bold(true),
		ModalBorder:   lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(c.accent),
		ModalTitle:    text.Copy().Bold(true),
		ModalHint:     dim.Copy(),
		ModalError:    lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5f5f")),
		InputText:     text.Copy(),
		InputHint:     dim.Copy(),
		InputCursor:   text.Copy(),
	}
}

package style

import "github.com/charmbracelet/lipgloss"

// ThemeNames is the ordered list of available theme names.
var ThemeNames = []string{"green", "dracula", "tokyo-night", "one-dark", "nord", "gruvbox", "monokai", "catppuccin"}

type Theme struct {
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
	SubjectTag    lipgloss.Style
}

type themeSpec struct {
	accent lipgloss.Color
	border lipgloss.Color
	rowBg  lipgloss.Color
}

var themeSpecs = map[string]themeSpec{
	// original
	"green": {accent: "#39ff14", border: "#1e5f1e", rowBg: "#0b140b"},
	// Dracula — pink keywords on dark slate
	"dracula": {accent: "#ff79c6", border: "#6272a4", rowBg: "#282a36"},
	// Tokyo Night — electric blue on deep navy
	"tokyo-night": {accent: "#7aa2f7", border: "#565f89", rowBg: "#1a1b2e"},
	// One Dark Pro — sky blue on charcoal
	"one-dark": {accent: "#61afef", border: "#5c6370", rowBg: "#21252b"},
	// Nord — arctic frost on polar night
	"nord": {accent: "#88c0d0", border: "#4c566a", rowBg: "#2e3440"},
	// Gruvbox — warm gold on dark brown
	"gruvbox": {accent: "#fabd2f", border: "#665c54", rowBg: "#1d2021"},
	// Monokai — vivid green on near-black
	"monokai": {accent: "#a6e22e", border: "#75715e", rowBg: "#272822"},
	// Catppuccin Mocha — lavender mauve on dark base
	"catppuccin": {accent: "#cba6f7", border: "#585b70", rowBg: "#181825"},
}

// ThemeOf returns the Theme for the given name, defaulting to "green" if unknown.
func ThemeOf(name string) Theme {
	spec, ok := themeSpecs[name]
	if !ok {
		spec = themeSpecs["green"]
	}
	return buildTheme(spec.accent, spec.border, spec.rowBg)
}

// NewTheme returns the default green theme (backwards compat).
func NewTheme() Theme {
	return ThemeOf("green")
}

func buildTheme(accent, border, rowBg lipgloss.Color) Theme {
	text := lipgloss.NewStyle().Foreground(accent)
	dim := lipgloss.NewStyle().Foreground(border)
	return Theme{
		Accent:        accent,
		Border:        border,
		Text:          text,
		Dim:           dim,
		Title:         text.Copy().Bold(true),
		TabActive:     lipgloss.NewStyle().Foreground(accent).Border(lipgloss.NormalBorder()).BorderForeground(accent).Padding(0, 1),
		TabInactive:   lipgloss.NewStyle().Foreground(border).Padding(1, 1),
		CheckboxDone:  dim.Copy(),
		CheckboxTodo:  text.Copy(),
		FooterHint:    dim.Copy(),
		ProjectDetail: dim.Copy(),
		SubjectActive: text.Copy().Bold(true),
		SubjectDim:    dim.Copy(),
		StatusDone:    lipgloss.NewStyle().Foreground(lipgloss.Color("#0c0c0c")).Background(lipgloss.Color("#3a7f3a")).Padding(0, 1),
		StatusInProg:  lipgloss.NewStyle().Foreground(lipgloss.Color("#0c0c0c")).Background(lipgloss.Color("#d4a017")).Padding(0, 1),
		StatusNotStr:  lipgloss.NewStyle().Foreground(lipgloss.Color("#0c0c0c")).Background(lipgloss.Color("#1e5f1e")).Padding(0, 1),
		RowActive:     lipgloss.NewStyle().Foreground(accent).Background(rowBg).Bold(true),
		ModalBorder:   lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(accent),
		ModalTitle:    text.Copy().Bold(true),
		ModalHint:     dim.Copy(),
		ModalError:    lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5f5f")),
		InputText:     text.Copy(),
		InputHint:     dim.Copy(),
		InputCursor:   text.Copy(),
		SubjectTag:    lipgloss.NewStyle().Foreground(accent).Background(border).Padding(0, 1).Bold(true),
	}
}

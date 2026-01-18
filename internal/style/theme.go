package style

import "github.com/charmbracelet/lipgloss"

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
}

func NewTheme() Theme {
	accent := lipgloss.Color("#39ff14")
	border := lipgloss.Color("#1e5f1e")
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
		RowActive:     lipgloss.NewStyle().Foreground(accent).Background(lipgloss.Color("#0b140b")).Bold(true),
		ModalBorder:   lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(accent),
		ModalTitle:    text.Copy().Bold(true),
		ModalHint:     dim.Copy(),
		ModalError:    lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5f5f")),
		InputText:     text.Copy(),
		InputHint:     dim.Copy(),
		InputCursor:   text.Copy(),
	}
}

package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"seman/internal/style"
	"seman/internal/ui/layout"
)

type ModalMode int

const (
	ModalHidden ModalMode = iota
	ModalForm
	ModalConfirm
)

type ModalField struct {
	Label string
	Value string
}

type ModalState struct {
	Mode    ModalMode
	Title   string
	Hint    string
	Message string
	Error   string
	Fields  []ModalField
}

func RenderModalArea(state ModalState, width, height int, t style.Theme) string {
	modal := RenderModal(state, width, t)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, modal)
}

func RenderModal(state ModalState, width int, t style.Theme) string {
	modalW := layout.MinInt(70, width-6)
	if modalW < 42 {
		modalW = 42
	}
	box := t.ModalBorder.Copy().Padding(1, 2).Width(modalW)

	var b strings.Builder
	title := t.ModalTitle.Render(state.Title)
	if title != "" {
		b.WriteString(title)
		b.WriteString("\n")
	}
	if state.Mode == ModalConfirm {
		b.WriteString("\n")
		b.WriteString(t.Text.Render(state.Message))
		b.WriteString("\n\n")
		b.WriteString(t.ModalHint.Render(state.Hint))
		return box.Render(b.String())
	}

	b.WriteString("\n")
	for i, field := range state.Fields {
		label := fmt.Sprintf("%-12s", field.Label+":")
		b.WriteString(t.Dim.Render(label))
		b.WriteString(" ")
		b.WriteString(field.Value)
		if i < len(state.Fields)-1 {
			b.WriteString("\n")
		}
	}
	if state.Error != "" {
		b.WriteString("\n\n")
		b.WriteString(t.ModalError.Render(state.Error))
	}
	if state.Hint != "" {
		b.WriteString("\n\n")
		b.WriteString(t.ModalHint.Render(state.Hint))
	}

	return box.Render(b.String())
}

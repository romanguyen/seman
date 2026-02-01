package components

import (
	"fmt"
	"strings"

	"seman/internal/domain"
	"seman/internal/style"
)

func RenderChecklist(items []domain.ChecklistItem, selected int, showCursor bool, t style.Theme) string {
	var b strings.Builder
	for i, item := range items {
		if i > 0 {
			b.WriteString("\n")
		}
		box := "[ ]"
		style := t.CheckboxTodo
		if item.Done {
			box = "[x]"
			style = t.CheckboxDone
		}
		if showCursor && i == selected {
			style = t.RowActive
		}
		line := fmt.Sprintf("%s %s", box, item.Text)
		b.WriteString(style.Render(line))
	}
	return b.String()
}

func RenderProjects(items []domain.ProjectItem, t style.Theme) string {
	var b strings.Builder
	for i, item := range items {
		if i > 0 {
			b.WriteString("\n")
		}
		name := t.Text.Render("- " + item.Name)
		meta := t.ProjectDetail.Render("  " + item.Subject + " - Due: " + item.Due)
		b.WriteString(name)
		b.WriteString("\n")
		b.WriteString(meta)
	}
	return b.String()
}

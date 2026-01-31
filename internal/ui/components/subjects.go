package components

import (
	"fmt"
	"strings"

	"seman/internal/domain"
	"seman/internal/style"
)

func RenderSubjects(items []domain.SubjectItem, selected int, t style.Theme) string {
	if len(items) == 0 {
		return t.Dim.Render("No subjects yet")
	}

	var b strings.Builder
	for i, subj := range items {
		if i > 0 {
			b.WriteString("\n\n")
		}
		prefix := "  "
		codeStyle := t.SubjectDim
		nameStyle := t.SubjectDim
		if i == selected {
			prefix = "> "
			codeStyle = t.SubjectActive
			nameStyle = t.SubjectActive
		}
		b.WriteString(codeStyle.Render(prefix + subj.Code))
		if subj.Name != "" {
			b.WriteString("\n")
			b.WriteString(nameStyle.Render("  " + subj.Name))
		}
	}
	return b.String()
}

func RenderSubjectTitle(items []domain.SubjectItem, selected int, t style.Theme) string {
	if selected < 0 || selected >= len(items) {
		return ""
	}
	subj := items[selected]
	title := fmt.Sprintf("%s (%s)", subj.Name, subj.Code)
	return t.Title.Render(title)
}

func RenderExamList(exams []domain.ExamItem, examIdx int, highlight bool, t style.Theme) string {
	if len(exams) == 0 {
		return t.Dim.Render("No exams this week")
	}

	var b strings.Builder
	for i, exam := range exams {
		if i > 0 {
			b.WriteString("\n\n")
		}
		nameStyle := t.Text
		prefix := "  "
		if highlight && i == examIdx {
			nameStyle = t.RowActive
			prefix = "> "
		}
		b.WriteString(nameStyle.Render(prefix + exam.Name))
		if exam.Date != "" {
			label := exam.Date
			if parsed, ok := domain.ParseExamDate(exam.Date); ok {
				label = domain.FormatDate(parsed)
			}
			b.WriteString("\n")
			b.WriteString(t.Dim.Render("Current date: " + label))
		}
		if len(exam.Retakes) > 0 {
			b.WriteString("\n")
			b.WriteString(t.Dim.Render("Retake dates:"))
			for _, date := range exam.Retakes {
				label := date
				if parsed, ok := domain.ParseExamDate(date); ok {
					label = domain.FormatDate(parsed)
				}
				b.WriteString("\n")
				b.WriteString(t.Dim.Render(" - " + label))
			}
		}
		if exam.Priority != "" {
			b.WriteString("\n")
			b.WriteString(t.Dim.Render("Priority: "))
			b.WriteString(RenderPriority(exam.Priority))
		}
	}

	return b.String()
}

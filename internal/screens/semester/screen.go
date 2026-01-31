package semester

import (
	"github.com/charmbracelet/lipgloss"
	"seman/internal/screens"
	"seman/internal/style"
	"seman/internal/ui/components"
	"seman/internal/ui/layout"
)

func Render(state screens.State, width, height int, t style.Theme) string {
	bounds := layout.ComputeSemesterLayout(width, height)

	leftBody := components.RenderSubjects(state.Subjects, state.SelectedSubj, t)
	leftPanel := components.RenderPanel(bounds.LeftWidth, height, "Subjects", leftBody, t)

	if bounds.RightWidth <= 0 {
		return leftPanel
	}

	subjectTitle := components.RenderSubjectTitle(state.Subjects, state.SelectedSubj, t)
	exams := components.RenderExamList(state.FilteredExams, state.ExamCursor, state.FocusExams, t)
	rightBody := subjectTitle + "\n" + exams
	rightPanel := components.RenderPanel(bounds.RightWidth, height, "", rightBody, t)

	return lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, " ", rightPanel)
}

package screens

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/style"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/ui/components"
)

func RenderSemester(state State, width, height int, t style.Theme) string {
	layout := components.ComputeSemesterLayout(width, height)

	leftBody := components.RenderSubjects(state.Subjects, state.SelectedSubj, t)
	leftPanel := components.RenderPanel(layout.LeftWidth, height, "Subjects", leftBody, t)

	if layout.RightWidth <= 0 {
		return leftPanel
	}

	subjectTitle := components.RenderSubjectTitle(state.Subjects, state.SelectedSubj, t)
	exams := components.RenderExamList(state.FilteredExams, state.ExamCursor, state.FocusExams, t)
	rightBody := subjectTitle + "\n" + exams
	rightPanel := components.RenderPanel(layout.RightWidth, height, "", rightBody, t)

	return lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, " ", rightPanel)
}

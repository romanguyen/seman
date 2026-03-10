package screens

import (
	"github.com/romanguyen/KEK-keep-everything-kool/internal/style"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/ui/components"
)

func RenderExams(state State, width, height int, t style.Theme) string {
	var header string
	if state.SubjectFilter != "" {
		header = t.Title.Render("Subject: "+state.SubjectFilter) + "  " + t.Dim.Render("[R] clear filter")
		header += "\n\n"
	}

	body := header + components.RenderFlatExams(state.FlatExams, state.ExamCursor, state.SubjectFilter, t)
	return components.RenderPanel(width, height, "Exams", body, t)
}

package projects

import (
	"student-exams-manager/internal/screens"
	"student-exams-manager/internal/style"
	"student-exams-manager/internal/ui/components"
	"student-exams-manager/internal/ui/layout"
)

func Render(state screens.State, width, height int, t style.Theme) string {
	bounds := layout.ComputeProjectsLayout(width, height)
	body := components.RenderProjectsTable(state.Projects, state.ProjectCursor, bounds.TableWidth, t)
	return components.RenderPanel(width, height, "Projects Overview", body, t)
}

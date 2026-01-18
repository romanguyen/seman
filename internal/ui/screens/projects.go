package screens

import (
	"student-exams-manager/internal/style"
	"student-exams-manager/internal/ui/components"
)

func RenderProjectsTab(state State, width, height int, t style.Theme) string {
	layout := components.ComputeProjectsLayout(width, height)
	body := components.RenderProjectsTable(state.Projects, state.ProjectCursor, layout.TableWidth, t)
	return components.RenderPanel(width, height, "Projects Overview", body, t)
}

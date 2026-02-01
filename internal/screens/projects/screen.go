package projects

import (
	"seman/internal/screens"
	"seman/internal/style"
	"seman/internal/ui/components"
	"seman/internal/ui/layout"
)

func Render(state screens.State, width, height int, t style.Theme) string {
	bounds := layout.ComputeProjectsLayout(width, height)
	body := components.RenderProjectsTable(state.Projects, state.ProjectCursor, bounds.TableWidth, t)
	return components.RenderPanel(width, height, "Projects Overview", body, t)
}

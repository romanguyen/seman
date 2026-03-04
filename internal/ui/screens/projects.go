package screens

import (
	"github.com/romanguyen/seman/internal/style"
	"github.com/romanguyen/seman/internal/ui/components"
)

func RenderProjectsTab(state State, width, height int, t style.Theme) string {
	layout := components.ComputeProjectsLayout(width, height)
	body := components.RenderProjectsTable(state.Projects, state.ProjectCursor, layout.TableWidth, t)
	return components.RenderPanel(width, height, "Projects Overview", body, t)
}

package screens

import (
	"github.com/romanguyen/KEK-keep-everything-kool/internal/style"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/ui/components"
)

func RenderProjectsTab(state State, width, height int, t style.Theme) string {
	layout := components.ComputeProjectsLayout(width, height)
	body := components.RenderProjectsTable(state.Projects, state.ProjectCursor, layout.TableWidth, t)
	return components.RenderPanel(width, height, "Projects Overview", body, t)
}

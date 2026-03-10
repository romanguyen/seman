package screens

import (
	"github.com/romanguyen/KEK-keep-everything-kool/internal/style"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/ui/components"
)

func RenderProjectsTab(state State, width, height int, t style.Theme) string {
	layout := components.ComputeProjectsLayout(width, height)

	title := "Projects Overview"
	var header string
	if state.SubjectFilter != "" {
		title = "Projects · " + state.SubjectFilter
		header = t.Title.Render("Subject: "+state.SubjectFilter) + "  " + t.Dim.Render("[R] clear filter") + "\n\n"
	}

	body := header + components.RenderProjectsTable(state.Projects, state.ProjectCursor, layout.TableWidth, t)
	return components.RenderPanel(width, height, title, body, t)
}

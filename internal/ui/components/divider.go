package components

import (
	"strings"

	"github.com/romanguyen/KEK-keep-everything-kool/internal/style"
)

func RenderDivider(width int, t style.Theme) string {
	if width < 1 {
		return ""
	}
	line := strings.Repeat("-", width)
	return t.Dim.Render(line)
}

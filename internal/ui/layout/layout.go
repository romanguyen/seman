package layout

const (
	HeaderHeight  = 3
	TabsHeight    = 3
	FooterHeight  = 3
	DividerHeight = 1
)

const (
	PanelPaddingX = 2
	PanelPaddingY = 1
	PanelBorderX  = 2
	PanelBorderY  = 2
	BarPaddingX   = 1
	BarPaddingY   = 0
	BarBorderX    = 2
)

type DashboardLayout struct {
	LeftWidth       int
	MiddleWidth     int
	RightWidth      int
	PanelHeight     int
	ChecklistWidth  int
	ChecklistHeight int
}

type SemesterLayout struct {
	LeftWidth  int
	RightWidth int
}

type TodoLayout struct {
	LeftWidth     int
	RightWidth    int
	PanelsHeight  int
	ActionsWidth  int
	ActionsHeight int
	SpacerHeight  int
}

type ProjectsLayout struct {
	TableWidth  int
	TableHeight int
}

func MainAreaHeight(totalHeight int) int {
	used := HeaderHeight + TabsHeight + DividerHeight + DividerHeight + FooterHeight
	mainHeight := totalHeight - used
	if mainHeight < 1 {
		return 1
	}
	return mainHeight
}

func PanelContentSize(width, height int) (int, int) {
	contentW := width - PanelBorderX - PanelPaddingX*2
	contentH := height - PanelBorderY - PanelPaddingY*2
	if contentW < 0 {
		contentW = 0
	}
	if contentH < 0 {
		contentH = 0
	}
	return contentW, contentH
}

func PanelContentWidth(width int) int {
	contentW := width - PanelBorderX - PanelPaddingX*2
	if contentW < 1 {
		contentW = 1
	}
	return contentW
}

func PanelHeightForLines(lines int) int {
	return lines + PanelPaddingY*2 + PanelBorderY
}

func ComputeDashboardLayout(width, height int) DashboardLayout {
	gap := 1
	available := width - gap*2
	if available < 0 {
		available = 0
	}

	colWidth := available / 3
	remainder := available % 3
	leftWidth := colWidth
	middleWidth := colWidth
	rightWidth := colWidth
	if remainder > 0 {
		leftWidth++
	}
	if remainder > 1 {
		middleWidth++
	}

	checkWidth, checkHeight := PanelContentSize(middleWidth, height)
	if checkHeight > 0 {
		checkHeight--
	}
	return DashboardLayout{
		LeftWidth:       leftWidth,
		MiddleWidth:     middleWidth,
		RightWidth:      rightWidth,
		PanelHeight:     height,
		ChecklistWidth:  checkWidth,
		ChecklistHeight: checkHeight,
	}
}

func ComputeSemesterLayout(width, height int) SemesterLayout {
	gap := 1
	available := width - gap
	if available < 0 {
		available = 0
	}
	leftWidth := available / 4
	rightWidth := available - leftWidth

	_ = height
	return SemesterLayout{
		LeftWidth:  leftWidth,
		RightWidth: rightWidth,
	}
}

func ComputeTodoLayout(width, height int) TodoLayout {
	available := width
	if available < 0 {
		available = 0
	}
	leftWidth := 0
	rightWidth := available

	spacerH := 0
	panelsH := height

	actionsW, actionsH := PanelContentSize(rightWidth, panelsH)
	if actionsH > 0 {
		actionsH--
	}

	return TodoLayout{
		LeftWidth:     leftWidth,
		RightWidth:    rightWidth,
		PanelsHeight:  panelsH,
		ActionsWidth:  actionsW,
		ActionsHeight: actionsH,
		SpacerHeight:  spacerH,
	}
}

func ComputeProjectsLayout(width, height int) ProjectsLayout {
	tableW, tableH := PanelContentSize(width, height)
	if tableH > 0 {
		tableH--
	}
	return ProjectsLayout{
		TableWidth:  tableW,
		TableHeight: tableH,
	}
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

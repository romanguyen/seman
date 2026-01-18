package components

const (
	headerHeight  = 3
	tabsHeight    = 3
	footerHeight  = 3
	dividerHeight = 1
)

const (
	panelPaddingX = 2
	panelPaddingY = 1
	panelBorderX  = 2
	panelBorderY  = 2
	barPaddingX   = 1
	barPaddingY   = 0
	barBorderX    = 2
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

type WeeklyLayout struct {
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
	used := headerHeight + tabsHeight + dividerHeight + dividerHeight + footerHeight
	mainHeight := totalHeight - used
	if mainHeight < 1 {
		return 1
	}
	return mainHeight
}

func PanelContentSize(width, height int) (int, int) {
	contentW := width - panelBorderX - panelPaddingX*2
	contentH := height - panelBorderY - panelPaddingY*2
	if contentW < 0 {
		contentW = 0
	}
	if contentH < 0 {
		contentH = 0
	}
	return contentW, contentH
}

func PanelContentWidth(width int) int {
	contentW := width - panelBorderX - panelPaddingX*2
	if contentW < 1 {
		contentW = 1
	}
	return contentW
}

func PanelHeightForLines(lines int) int {
	return lines + panelPaddingY*2 + panelBorderY
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

func ComputeWeeklyLayout(width, height int) WeeklyLayout {
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

	return WeeklyLayout{
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

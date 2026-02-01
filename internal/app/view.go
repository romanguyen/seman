package app

import (
	"strings"

	"seman/internal/screens"
	"seman/internal/screens/dashboard"
	"seman/internal/screens/lofi"
	"seman/internal/screens/projects"
	"seman/internal/screens/semester"
	"seman/internal/screens/settings"
	"seman/internal/screens/todo"
	"seman/internal/style"
	"seman/internal/ui/components"
	"seman/internal/ui/layout"
)

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Loading..."
	}

	t := style.NewThemeWithName(m.themeName)
	header := components.RenderHeader(m.width, t)
	tabs := components.RenderTabs(m.activeTab, m.width, m.weekLabel, m.tabItems(), t)
	divider := components.RenderDivider(m.width, t)
	mainHeight := layout.MainAreaHeight(m.height)
	state := m.viewState()
	var main string
	switch state.ActiveTab {
	case tabDashboard:
		main = dashboard.Render(state, m.width, mainHeight, t)
	case tabExams:
		main = semester.Render(state, m.width, mainHeight, t)
	case tabTodos:
		main = todo.Render(state, m.width, mainHeight, t)
	case tabProjects:
		main = projects.Render(state, m.width, mainHeight, t)
	case tabSettings:
		main = settings.Render(state, m.width, mainHeight, t)
	case tabLofi:
		main = lofi.Render(state, m.width, mainHeight, t)
	default:
		main = screens.RenderPlaceholder(m.width, mainHeight, t)
	}
	if state.Modal.Mode != components.ModalHidden {
		main = screens.RenderModal(state, m.width, mainHeight, t)
	}
	footer := components.RenderFooter(m.width, len(m.tabItems()), m.activeTab, m.saveError, t)

	return strings.Join([]string{header, tabs, divider, main, divider, footer}, "\n")
}

func (m Model) viewState() screens.State {
	start, end, all := m.weekRange()
	state := screens.State{
		ActiveTab:     m.activeTab,
		ConfirmOn:     m.confirmOn,
		ThemeName:     m.themeName,
		Checklist:     m.checklist,
		Projects:      m.projects,
		Subjects:      m.subjects,
		SelectedSubj:  m.selectedSubj,
		ExamCursor:    m.visibleIndex(m.examVisible, m.examCursor),
		FilteredExams: m.filteredExams(),
		FocusExams:    m.semesterFocus == focusExams,
		WeekLabel:     m.weekLabel,
		FilterStart:   start,
		FilterEnd:     end,
		FilterAll:     all,
		WeekSpan:      m.weekSpan,
		ProjectCursor: m.projectCursor,
		LofiEnabled:   m.lofi.enabled,
		LofiURL:       m.lofi.url,
		LofiStatus:    m.lofi.status,
		LofiError:     m.lofi.err,
		LofiPlaylist:  m.lofiPlaylist,
		LofiCursor:    m.lofiCursor,
		LofiNow:       m.lofiNow,
	}

	if m.modal == modalNone {
		return state
	}

	fields := make([]components.ModalField, 0, len(m.formFields))
	for _, field := range m.formFields {
		fields = append(fields, components.ModalField{
			Label: field.label,
			Value: field.input.View(),
		})
	}

	modalState := components.ModalState{
		Title:  m.modalTitle,
		Hint:   m.modalHint,
		Fields: fields,
	}
	switch m.modal {
	case modalConfirm:
		modalState.Mode = components.ModalConfirm
		modalState.Message = m.modalError
	default:
		modalState.Mode = components.ModalForm
		modalState.Error = m.modalError
	}
	state.Modal = modalState
	return state
}

func (m Model) tabItems() []components.TabItem {
	items := []components.TabItem{
		{ID: tabDashboard, Label: "Dashboard"},
		{ID: tabExams, Label: "Exams"},
		{ID: tabTodos, Label: "Todos"},
		{ID: tabProjects, Label: "Projects"},
		{ID: tabSettings, Label: "Settings"},
	}
	if m.lofi.enabled {
		items = append(items, components.TabItem{ID: tabLofi, Label: "Lofi"})
	}
	return items
}

func (m *Model) switchToTab(index int) bool {
	items := m.tabItems()
	if index < 0 || index >= len(items) {
		return false
	}
	m.activeTab = items[index].ID
	m.resize(m.width, m.height)
	return true
}

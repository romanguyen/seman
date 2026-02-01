package app

import (
	"strings"

	"seman/internal/domain"
	"seman/internal/storage"
	"seman/internal/style"
)

func (m *Model) applyData(data storage.SemesterData) {
	m.subjects = data.Subjects
	m.projects = data.Projects
	m.checklistItems = data.Checklist
	m.confirmOn = data.ConfirmOn
	m.setWeekSpanFromData(data.WeekSpan)
	m.setWeekStartFromData(data.WeekStart)
	m.lofi.enabled = data.LofiEnabled
	m.lofi.url = strings.TrimSpace(data.LofiURL)
	m.lofi.status = lofiStatusStopped
	m.lofi.err = ""
	m.setThemeFromData(data.ThemeName)
	m.ensureTodoDueDates()
	m.refreshExamFilter()
	m.refreshChecklistView()

	m.selectedSubj = clampIndex(m.selectedSubj, len(m.subjects))
	m.projectCursor = clampIndex(m.projectCursor, len(m.projects))
	m.lofiNow = clampIndex(m.lofiNow, len(m.lofiPlaylist))
	m.lofiCursor = clampIndex(m.lofiCursor, len(m.lofiPlaylist))
}

func (m Model) exportData() storage.SemesterData {
	return storage.SemesterData{
		Subjects:    m.subjects,
		Projects:    m.projects,
		Checklist:   m.checklistItems,
		ConfirmOn:   m.confirmOn,
		WeekStart:   domain.FormatDate(m.weekStart),
		WeekSpan:    m.weekSpan,
		LofiEnabled: m.lofi.enabled,
		LofiURL:     m.lofi.url,
		ThemeName:   string(m.themeName),
	}
}

func (m *Model) setThemeFromData(name string) {
	switch style.ThemeName(name) {
	case style.ThemeDracula, style.ThemeNord, style.ThemeSolarize, style.ThemeCyberpnk:
		m.themeName = style.ThemeName(name)
	default:
		m.themeName = style.ThemeMatrix
	}
}

func (m *Model) cycleTheme() {
	m.themeName = style.NextTheme(m.themeName)
	m.persist()
}

func (m *Model) persist() {
	if m.store == nil {
		return
	}
	if err := m.store.Save(m.exportData()); err != nil {
		m.saveError = "Save failed: " + err.Error()
		return
	}
	m.saveError = ""
}

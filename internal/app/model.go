package app

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"seman/internal/domain"
	"seman/internal/storage"
)

type Model struct {
	width           int
	height          int
	activeTab       int
	checklist       viewport.Model
	checklistItems  []domain.ChecklistItem
	checklistCursor int
	todoVisible     []int
	projects        []domain.ProjectItem
	subjects        []domain.SubjectItem
	selectedSubj    int
	examCursor      int
	examVisible     []int
	semesterFocus   semesterFocus
	weekLabel       string
	weekStart       time.Time
	projectCursor   int
	confirmOn       bool
	weekSpan        int
	lofi            lofiState
	lofiPlaylist    []domain.LofiTrack
	lofiCursor      int
	lofiNow         int
	lofiReload      bool
	modal           modalKind
	formFields      []formField
	formFocus       int
	modalTitle      string
	modalHint       string
	modalError      string
	confirmAction   confirmAction
	editSubjectIdx  int
	editExamIdx     int
	editProjectIdx  int
	editTodoIdx     int
	saveError       string
	store           storage.Store
}

type semesterFocus int

const (
	focusSubjects semesterFocus = iota
	focusExams
)

const (
	tabDashboard = iota
	tabExams
	tabTodos
	tabProjects
	tabSettings
	tabLofi
)

func NewModel(store storage.Store, data storage.SemesterData, hasData bool) Model {
	if !hasData {
		data = DefaultData()
	}
	vp := viewport.New(0, 0)
	m := Model{
		activeTab:      tabDashboard,
		checklist:      vp,
		selectedSubj:   0,
		examCursor:     0,
		semesterFocus:  focusSubjects,
		projectCursor:  0,
		editSubjectIdx: -1,
		editExamIdx:    -1,
		editProjectIdx: -1,
		editTodoIdx:    -1,
		lofiNow:        -1,
		store:          store,
	}
	m.applyData(data)
	return m
}

func (m Model) Init() tea.Cmd {
	if m.lofi.enabled && strings.TrimSpace(m.lofi.url) != "" {
		return loadLofiPlaylist(m.lofi.url)
	}
	return nil
}

func (m *Model) openAddForTab() bool {
	switch m.activeTab {
	case tabExams:
		m.openAddExam()
		return true
	case tabTodos:
		m.openAddTodo()
		return true
	case tabProjects:
		m.openAddProject()
		return true
	}
	return false
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.resize(msg.Width, msg.Height)
		return m, nil
	case lofiExitMsg:
		m.handleLofiExit(msg)
		return m, nil
	case lofiPlaylistMsg:
		m.applyLofiPlaylist(msg)
		return m, nil
	case lofiIndexMsg:
		cmd := m.applyLofiIndex(msg)
		return m, cmd
	case lofiPlaybackMsg:
		cmd := m.applyLofiPlayback(msg)
		return m, cmd
	case tea.KeyMsg:
		key := msg.String()
		if m.modal != modalNone {
			return m.updateModal(msg)
		}
		if len(key) == 1 && key[0] >= '1' && key[0] <= '9' {
			if m.switchToTab(int(key[0] - '1')) {
				return m, nil
			}
		}
		switch key {
		case "ctrl+c", "q":
			m.shutdownLofi()
			return m, tea.Quit
		case "left":
			m.shiftWeek(-1)
			return m, nil
		case "right":
			m.shiftWeek(1)
			return m, nil
		case "a", "A":
			if m.openAddForTab() {
				return m, nil
			}
		case "s", "S":
			if m.activeTab == tabExams {
				m.openAddSubject()
				return m, nil
			}
		case "e", "E":
			m.openEditCurrent()
			return m, nil
		case "n", "N":
			if m.activeTab == tabTodos {
				m.openAddTodo()
				return m, nil
			}
		case "d", "D":
			m.queueDelete()
			return m, nil
		case "c", "C":
			if m.activeTab == tabSettings {
				m.queueClearAll()
				return m, nil
			}
		case "o", "O":
			if m.activeTab == tabSettings {
				m.confirmOn = !m.confirmOn
				m.persist()
				return m, nil
			}
		case "w", "W":
			if m.activeTab == tabSettings {
				m.cycleWeekSpan()
				return m, nil
			}
		case "l", "L":
			if m.activeTab == tabSettings {
				cmd := m.toggleLofiEnabled()
				return m, cmd
			}
		case "u", "U":
			if m.activeTab == tabSettings || m.activeTab == tabLofi {
				m.openEditLofiURL()
				return m, nil
			}
		}

		if m.activeTab == tabTodos {
			switch key {
			case "j", "down":
				m.moveChecklistCursor(1)
				return m, nil
			case "k", "up":
				m.moveChecklistCursor(-1)
				return m, nil
			case "pgdown":
				m.moveChecklistCursor(m.checklist.Height)
				return m, nil
			case "pgup":
				m.moveChecklistCursor(-m.checklist.Height)
				return m, nil
			case " ", "enter", "x", "X":
				m.toggleChecklistItem()
				return m, nil
			}

			var cmd tea.Cmd
			m.checklist, cmd = m.checklist.Update(msg)
			return m, cmd
		}

		if m.activeTab == tabDashboard {
			switch key {
			case "j", "down":
				m.checklist.LineDown(1)
				return m, nil
			case "k", "up":
				m.checklist.LineUp(1)
				return m, nil
			case "pgdown":
				m.checklist.ViewDown()
				return m, nil
			case "pgup":
				m.checklist.ViewUp()
				return m, nil
			}

			var cmd tea.Cmd
			m.checklist, cmd = m.checklist.Update(msg)
			return m, cmd
		}

		if m.activeTab == tabExams {
			switch key {
			case "tab":
				m.toggleSemesterFocus()
				return m, nil
			case "j", "down":
				m.moveSemesterCursor(1)
				return m, nil
			case "k", "up":
				m.moveSemesterCursor(-1)
				return m, nil
			}
		}

		if m.activeTab == tabProjects {
			switch key {
			case "j", "down":
				if m.projectCursor < len(m.projects)-1 {
					m.projectCursor++
				}
				return m, nil
			case "k", "up":
				if m.projectCursor > 0 {
					m.projectCursor--
				}
				return m, nil
			}
		}

		if m.activeTab == tabLofi {
			switch key {
			case "enter":
				cmd := m.playLofiAt(m.lofiCursor)
				return m, cmd
			case " ":
				cmd := m.toggleLofiPlayPause()
				return m, cmd
			case "n", "N":
				cmd := m.lofiNext()
				return m, cmd
			case "b", "B":
				cmd := m.lofiPrev()
				return m, cmd
			case "x", "X":
				m.lofiStop()
				return m, nil
			}
		}
	}

	return m, nil
}

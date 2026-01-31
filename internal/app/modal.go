package app

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"student-exams-manager/internal/domain"
	"student-exams-manager/internal/style"
	"student-exams-manager/internal/ui/layout"
)

type modalKind int

const (
	modalNone modalKind = iota
	modalAddSubject
	modalAddExam
	modalAddProject
	modalAddTodo
	modalEditSubject
	modalEditExam
	modalEditProject
	modalEditTodo
	modalEditLofiURL
	modalConfirm
)

type confirmKind int

const (
	confirmDeleteSubject confirmKind = iota
	confirmDeleteProject
	confirmDeleteTodo
	confirmClearAll
)

type confirmAction struct {
	kind       confirmKind
	subjectIdx int
	projectIdx int
}

type formField struct {
	label    string
	input    textinput.Model
	required bool
}

type formSpec struct {
	label    string
	required bool
}

var (
	subjectFormSpec = []formSpec{
		{label: "Code", required: true},
		{label: "Name", required: true},
	}
	examFormSpec = []formSpec{
		{label: "Subject", required: true},
		{label: "Exam Name", required: true},
		{label: "Date", required: true},
		{label: "Retakes", required: false},
		{label: "Priority", required: false},
	}
	projectFormSpec = []formSpec{
		{label: "Name", required: true},
		{label: "Subject", required: true},
		{label: "Deadline", required: true},
		{label: "Status", required: false},
	}
	todoFormSpec = []formSpec{
		{label: "Task", required: true},
	}
	lofiFormSpec = []formSpec{
		{label: "Playlist URL", required: true},
	}
)

func (m Model) updateModal(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "esc":
			m.closeModal()
			return m, nil
		case "enter":
			if m.modal == modalConfirm {
				m.applyConfirmAction()
				m.closeModal()
				return m, nil
			}
			if m.formFocus == len(m.formFields)-1 {
				if err := m.submitForm(); err != nil {
					m.modalError = err.Error()
					return m, nil
				}
				cmd := m.consumeLofiReload()
				m.closeModal()
				return m, cmd
			}
			m.setFormFocus(m.formFocus + 1)
			return m, nil
		case "tab":
			if m.modal != modalConfirm {
				m.setFormFocus((m.formFocus + 1) % len(m.formFields))
				return m, nil
			}
		case "shift+tab":
			if m.modal != modalConfirm {
				next := m.formFocus - 1
				if next < 0 {
					next = len(m.formFields) - 1
				}
				m.setFormFocus(next)
				return m, nil
			}
		case "y", "Y":
			if m.modal == modalConfirm {
				m.applyConfirmAction()
				m.closeModal()
				return m, nil
			}
		case "n", "N":
			if m.modal == modalConfirm {
				m.closeModal()
				return m, nil
			}
		}
	}

	if m.modal == modalConfirm || len(m.formFields) == 0 {
		return m, nil
	}

	field := m.formFields[m.formFocus]
	var cmd tea.Cmd
	field.input, cmd = field.input.Update(msg)
	m.formFields[m.formFocus] = field
	return m, cmd
}

func (m *Model) closeModal() {
	m.modal = modalNone
	m.formFields = nil
	m.formFocus = 0
	m.modalTitle = ""
	m.modalHint = ""
	m.modalError = ""
	m.editSubjectIdx = -1
	m.editExamIdx = -1
	m.editProjectIdx = -1
	m.editTodoIdx = -1
}

func (m *Model) setFormFocus(idx int) {
	if len(m.formFields) == 0 {
		return
	}
	for i := range m.formFields {
		if i == idx {
			m.formFields[i].input.Focus()
		} else {
			m.formFields[i].input.Blur()
		}
	}
	m.formFocus = idx
}

func buildForm(specs []formSpec, width int) []formField {
	fields := make([]formField, 0, len(specs))
	for _, spec := range specs {
		fields = append(fields, newFormField(spec.label, width, spec.required))
	}
	return fields
}

func setFormValues(fields []formField, values ...string) {
	for i := 0; i < len(fields) && i < len(values); i++ {
		fields[i].input.SetValue(values[i])
	}
}

func (m *Model) openAddSubject() {
	fields := buildForm(subjectFormSpec, m.modalInputWidth())
	m.openFormModal(modalAddSubject, "Add Subject", fields)
}

func (m *Model) openAddExam() {
	fields := buildForm(examFormSpec, m.modalInputWidth())
	m.openFormModal(modalAddExam, "Add Exam", fields)
}

func (m *Model) openAddProject() {
	fields := buildForm(projectFormSpec, m.modalInputWidth())
	m.openFormModal(modalAddProject, "Add Project", fields)
}

func (m *Model) openAddTodo() {
	fields := buildForm(todoFormSpec, m.modalInputWidth())
	m.openFormModal(modalAddTodo, "Add Todo", fields)
}

func (m *Model) openEditCurrent() {
	switch m.activeTab {
	case tabExams:
		if m.semesterFocus == focusExams {
			m.openEditExam()
			return
		}
		m.openEditSubject()
	case tabTodos:
		m.openEditTodo()
	case tabProjects:
		m.openEditProject()
	}
}

func (m *Model) openEditSubject() {
	if m.selectedSubj < 0 || m.selectedSubj >= len(m.subjects) {
		return
	}
	subj := m.subjects[m.selectedSubj]
	fields := buildForm(subjectFormSpec, m.modalInputWidth())
	setFormValues(fields, subj.Code, subj.Name)
	m.editSubjectIdx = m.selectedSubj
	m.openFormModal(modalEditSubject, "Edit Subject", fields)
}

func (m *Model) openEditExam() {
	exams := m.examsForSelected()
	if len(exams) == 0 || m.examCursor < 0 || m.examCursor >= len(exams) {
		return
	}
	exam := exams[m.examCursor]
	fields := buildForm(examFormSpec[1:], m.modalInputWidth())
	setFormValues(fields, exam.Name, exam.Date, strings.Join(exam.Retakes, ", "), exam.Priority)
	m.editSubjectIdx = m.selectedSubj
	m.editExamIdx = m.examCursor
	m.openFormModal(modalEditExam, "Edit Exam", fields)
}

func (m *Model) openEditProject() {
	if m.projectCursor < 0 || m.projectCursor >= len(m.projects) {
		return
	}
	project := m.projects[m.projectCursor]
	fields := buildForm(projectFormSpec, m.modalInputWidth())
	setFormValues(fields, project.Name, project.Subject, project.Due, project.Status)
	m.editProjectIdx = m.projectCursor
	m.openFormModal(modalEditProject, "Edit Project", fields)
}

func (m *Model) openEditTodo() {
	if m.checklistCursor < 0 || m.checklistCursor >= len(m.checklistItems) {
		return
	}
	item := m.checklistItems[m.checklistCursor]
	fields := buildForm(todoFormSpec, m.modalInputWidth())
	setFormValues(fields, item.Text)
	m.editTodoIdx = m.checklistCursor
	m.openFormModal(modalEditTodo, "Edit Todo", fields)
}

func (m *Model) openEditLofiURL() {
	fields := buildForm(lofiFormSpec, m.modalInputWidth())
	setFormValues(fields, m.lofi.url)
	m.openFormModal(modalEditLofiURL, "Edit Lofi Playlist", fields)
}

func (m *Model) openFormModal(kind modalKind, title string, fields []formField) {
	m.modal = kind
	m.formFields = fields
	m.modalTitle = title
	m.modalHint = "Tab to switch, Enter to save, Esc to cancel"
	m.modalError = ""
	m.setFormFocus(0)
}

func newFormField(label string, width int, required bool) formField {
	input := textinput.New()
	input.Prompt = ""
	input.Width = width
	input.CharLimit = 200
	t := style.NewTheme()
	input.TextStyle = t.InputText
	input.PlaceholderStyle = t.InputHint
	input.CursorStyle = t.InputCursor
	return formField{
		label:    label,
		input:    input,
		required: required,
	}
}

func (m *Model) formValue(idx int) string {
	if idx < 0 || idx >= len(m.formFields) {
		return ""
	}
	return strings.TrimSpace(m.formFields[idx].input.Value())
}

func (m *Model) saveSubject(idx int, code, name string) error {
	if code == "" || name == "" {
		return fmt.Errorf("Code and Name are required.")
	}
	if idx < 0 {
		m.subjects = append(m.subjects, domain.SubjectItem{Code: code, Name: name})
		m.selectedSubj = len(m.subjects) - 1
		m.persist()
		return nil
	}
	if idx >= len(m.subjects) {
		return nil
	}
	m.subjects[idx].Code = code
	m.subjects[idx].Name = name
	m.persist()
	return nil
}

func (m *Model) addExam(subjectCode, examName, date, retakesRaw, priority string) error {
	if subjectCode == "" || examName == "" || date == "" {
		return fmt.Errorf("Subject, Exam Name, and Date are required.")
	}
	idx := findSubjectIndex(m.subjects, subjectCode)
	if idx < 0 {
		return fmt.Errorf("Subject code not found.")
	}
	retakes := splitCSV(retakesRaw)
	m.subjects[idx].Exams = append(m.subjects[idx].Exams, domain.ExamItem{
		Name:     examName,
		Date:     date,
		Retakes:  retakes,
		Priority: strings.ToUpper(priority),
	})
	m.selectedSubj = idx
	m.examCursor = len(m.subjects[idx].Exams) - 1
	m.refreshExamFilter()
	m.persist()
	return nil
}

func (m *Model) updateExam(subjectIdx, examIdx int, examName, date, retakesRaw, priority string) error {
	if subjectIdx < 0 || subjectIdx >= len(m.subjects) {
		return nil
	}
	exams := m.subjects[subjectIdx].Exams
	if examIdx < 0 || examIdx >= len(exams) {
		return nil
	}
	if examName == "" || date == "" {
		return fmt.Errorf("Exam Name and Date are required.")
	}
	exams[examIdx].Name = examName
	exams[examIdx].Date = date
	exams[examIdx].Retakes = splitCSV(retakesRaw)
	exams[examIdx].Priority = strings.ToUpper(priority)
	m.subjects[subjectIdx].Exams = exams
	m.examCursor = examIdx
	m.refreshExamFilter()
	m.persist()
	return nil
}

func (m *Model) saveProject(idx int, name, subject, deadline, status string) error {
	if name == "" || subject == "" || deadline == "" {
		return fmt.Errorf("Name, Subject, and Deadline are required.")
	}
	if status == "" {
		status = domain.ProjectStatusNotStarted
	}
	status = strings.ToUpper(status)
	if idx < 0 {
		m.projects = append(m.projects, domain.ProjectItem{
			Name:    name,
			Subject: subject,
			Due:     deadline,
			Status:  status,
		})
		m.projectCursor = len(m.projects) - 1
		m.persist()
		return nil
	}
	if idx >= len(m.projects) {
		return nil
	}
	m.projects[idx].Name = name
	m.projects[idx].Subject = subject
	m.projects[idx].Due = deadline
	m.projects[idx].Status = status
	m.projectCursor = idx
	m.persist()
	return nil
}

func (m *Model) saveTodo(idx int, task string) error {
	if task == "" {
		return fmt.Errorf("Task is required.")
	}
	if idx < 0 {
		m.checklistItems = append(m.checklistItems, domain.ChecklistItem{
			Text: task,
			Done: false,
			Due:  m.weekStart.Format("2006-01-02"),
		})
		m.checklistCursor = len(m.checklistItems) - 1
		m.persist()
		m.refreshChecklistView()
		return nil
	}
	if idx >= len(m.checklistItems) {
		return nil
	}
	m.checklistItems[idx].Text = task
	m.persist()
	m.refreshChecklistView()
	return nil
}

func (m *Model) submitForm() error {
	switch m.modal {
	case modalAddSubject:
		return m.saveSubject(-1, m.formValue(0), m.formValue(1))
	case modalAddExam:
		return m.addExam(m.formValue(0), m.formValue(1), m.formValue(2), m.formValue(3), m.formValue(4))
	case modalAddProject:
		return m.saveProject(-1, m.formValue(0), m.formValue(1), m.formValue(2), m.formValue(3))
	case modalAddTodo:
		return m.saveTodo(-1, m.formValue(0))
	case modalEditSubject:
		return m.saveSubject(m.editSubjectIdx, m.formValue(0), m.formValue(1))
	case modalEditExam:
		return m.updateExam(m.editSubjectIdx, m.editExamIdx, m.formValue(0), m.formValue(1), m.formValue(2), m.formValue(3))
	case modalEditProject:
		return m.saveProject(m.editProjectIdx, m.formValue(0), m.formValue(1), m.formValue(2), m.formValue(3))
	case modalEditTodo:
		return m.saveTodo(m.editTodoIdx, m.formValue(0))
	case modalEditLofiURL:
		url := m.formValue(0)
		if url == "" {
			return fmt.Errorf("Playlist URL is required.")
		}
		m.lofi.url = url
		m.lofiReload = true
		m.persist()
		if m.lofi.cmd != nil && m.lofi.status != lofiStatusStopped {
			_ = m.sendLofiCommand("loadfile", m.lofi.url, "replace")
			m.lofi.status = lofiStatusPlaying
		}
	}
	return nil
}

func findSubjectIndex(items []domain.SubjectItem, code string) int {
	code = strings.ToUpper(strings.TrimSpace(code))
	for i, item := range items {
		if strings.ToUpper(item.Code) == code {
			return i
		}
	}
	return -1
}

func splitCSV(raw string) []string {
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		val := strings.TrimSpace(part)
		if val != "" {
			out = append(out, val)
		}
	}
	return out
}

func (m *Model) queueDelete() {
	switch m.activeTab {
	case tabExams:
		if len(m.subjects) == 0 {
			return
		}
		action := confirmAction{kind: confirmDeleteSubject, subjectIdx: m.selectedSubj}
		message := fmt.Sprintf("Delete subject %s and its exams?", m.subjects[m.selectedSubj].Code)
		m.confirmOrApply(action, message)
	case tabProjects:
		if len(m.projects) == 0 {
			return
		}
		action := confirmAction{kind: confirmDeleteProject, projectIdx: m.projectCursor}
		message := fmt.Sprintf("Delete project %s?", m.projects[m.projectCursor].Name)
		m.confirmOrApply(action, message)
	case tabTodos:
		if len(m.checklistItems) == 0 || m.checklistCursor < 0 || m.checklistCursor >= len(m.checklistItems) {
			return
		}
		action := confirmAction{kind: confirmDeleteTodo, projectIdx: m.checklistCursor}
		message := fmt.Sprintf("Delete task \"%s\"?", m.checklistItems[m.checklistCursor].Text)
		m.confirmOrApply(action, message)
	}
}

func (m *Model) queueClearAll() {
	action := confirmAction{kind: confirmClearAll}
	message := "Clear all data? This cannot be undone."
	m.confirmOrApply(action, message)
}

func (m *Model) confirmOrApply(action confirmAction, message string) {
	if m.confirmOn {
		m.modal = modalConfirm
		m.confirmAction = action
		m.modalTitle = "Confirm"
		m.modalHint = "[Y] Confirm  [N] Cancel"
		m.modalError = message
		return
	}
	m.confirmAction = action
	m.applyConfirmAction()
}

func (m *Model) applyConfirmAction() {
	switch m.confirmAction.kind {
	case confirmDeleteSubject:
		if m.confirmAction.subjectIdx >= 0 && m.confirmAction.subjectIdx < len(m.subjects) {
			m.subjects = append(m.subjects[:m.confirmAction.subjectIdx], m.subjects[m.confirmAction.subjectIdx+1:]...)
			if m.selectedSubj >= len(m.subjects) {
				m.selectedSubj = len(m.subjects) - 1
			}
			if m.selectedSubj < 0 {
				m.selectedSubj = 0
			}
		}
	case confirmDeleteProject:
		if m.confirmAction.projectIdx >= 0 && m.confirmAction.projectIdx < len(m.projects) {
			m.projects = append(m.projects[:m.confirmAction.projectIdx], m.projects[m.confirmAction.projectIdx+1:]...)
			if m.projectCursor >= len(m.projects) {
				m.projectCursor = len(m.projects) - 1
			}
			if m.projectCursor < 0 {
				m.projectCursor = 0
			}
		}
	case confirmDeleteTodo:
		if m.confirmAction.projectIdx >= 0 && m.confirmAction.projectIdx < len(m.checklistItems) {
			idx := m.confirmAction.projectIdx
			m.checklistItems = append(m.checklistItems[:idx], m.checklistItems[idx+1:]...)
			if m.checklistCursor >= len(m.checklistItems) {
				m.checklistCursor = len(m.checklistItems) - 1
			}
			if m.checklistCursor < 0 {
				m.checklistCursor = 0
			}
			m.refreshChecklistView()
		}
	case confirmClearAll:
		m.subjects = nil
		m.projects = nil
		m.checklistItems = nil
		m.weeklyExams = nil
		m.selectedSubj = 0
		m.projectCursor = 0
		m.refreshChecklistView()
	}
	m.persist()
}

func (m Model) modalInputWidth() int {
	modalW := layout.MinInt(70, m.width-6)
	if modalW < 42 {
		modalW = 42
	}
	return layout.MaxInt(12, modalW-18)
}

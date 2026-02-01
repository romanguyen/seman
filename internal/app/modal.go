package app

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"seman/internal/domain"
	"seman/internal/style"
	"seman/internal/ui/layout"
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
		{label: "Date (DD/MM/YYYY)", required: true},
		{label: "Retakes (DD/MM/YYYY, ...)", required: false},
		{label: "Priority", required: false},
	}
	projectFormSpec = []formSpec{
		{label: "Name", required: true},
		{label: "Subject", required: true},
		{label: "Deadline (DD/MM/YYYY)", required: true},
		{label: "Status", required: false},
	}
	todoFormSpec = []formSpec{
		{label: "Task", required: true},
		{label: "Due (DD/MM/YYYY)", required: true},
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

func (m *Model) openForm(kind modalKind, title string, specs []formSpec, values ...string) {
	fields := buildForm(specs, m.modalInputWidth())
	setFormValues(fields, values...)
	m.openFormModal(kind, title, fields)
}

func formatDateForInput(value string) string {
	if parsed, ok := domain.ParseExamDate(value); ok {
		return domain.FormatDate(parsed)
	}
	return strings.TrimSpace(value)
}

func (m *Model) openAddSubject() {
	m.openForm(modalAddSubject, "Add Subject", subjectFormSpec)
}

func (m *Model) openAddExam() {
	m.openForm(modalAddExam, "Add Exam", examFormSpec)
}

func (m *Model) openAddProject() {
	m.openForm(modalAddProject, "Add Project", projectFormSpec)
}

func (m *Model) openAddTodo() {
	m.openForm(modalAddTodo, "Add Todo", todoFormSpec, "", domain.FormatDate(m.weekStart))
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
	if !inBounds(m.selectedSubj, len(m.subjects)) {
		return
	}
	subj := m.subjects[m.selectedSubj]
	m.editSubjectIdx = m.selectedSubj
	m.openForm(modalEditSubject, "Edit Subject", subjectFormSpec, subj.Code, subj.Name)
}

func (m *Model) openEditExam() {
	exams := m.examsForSelected()
	if !inBounds(m.examCursor, len(exams)) {
		return
	}
	exam := exams[m.examCursor]
	retakes := make([]string, 0, len(exam.Retakes))
	for _, date := range exam.Retakes {
		retakes = append(retakes, formatDateForInput(date))
	}
	m.editSubjectIdx = m.selectedSubj
	m.editExamIdx = m.examCursor
	m.openForm(modalEditExam, "Edit Exam", examFormSpec[1:], exam.Name, formatDateForInput(exam.Date), strings.Join(retakes, ", "), exam.Priority)
}

func (m *Model) openEditProject() {
	if !inBounds(m.projectCursor, len(m.projects)) {
		return
	}
	project := m.projects[m.projectCursor]
	m.editProjectIdx = m.projectCursor
	m.openForm(modalEditProject, "Edit Project", projectFormSpec, project.Name, project.Subject, formatDateForInput(project.Due), project.Status)
}

func (m *Model) openEditTodo() {
	if !inBounds(m.checklistCursor, len(m.checklistItems)) {
		return
	}
	item := m.checklistItems[m.checklistCursor]
	m.editTodoIdx = m.checklistCursor
	m.openForm(modalEditTodo, "Edit Todo", todoFormSpec, item.Text, formatDateForInput(item.Due))
}

func (m *Model) openEditLofiURL() {
	m.openForm(modalEditLofiURL, "Edit Lofi Playlist", lofiFormSpec, m.lofi.url)
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
	input.Cursor.Style = t.InputCursor
	return formField{
		label:    label,
		input:    input,
		required: required,
	}
}

func (m *Model) formValue(idx int) string {
	if !inBounds(idx, len(m.formFields)) {
		return ""
	}
	return strings.TrimSpace(m.formFields[idx].input.Value())
}

func parseStrictDateOrError(value, message string) (string, error) {
	parsed, ok := domain.ParseStrictDate(value)
	if !ok {
		return "", errors.New(message)
	}
	return domain.FormatDate(parsed), nil
}

func parseRetakesOrError(raw string) ([]string, error) {
	retakes := splitCSV(raw)
	for i := range retakes {
		parsed, ok := domain.ParseStrictDate(retakes[i])
		if !ok {
			return nil, fmt.Errorf("retakes must be DD/MM/YYYY")
		}
		retakes[i] = domain.FormatDate(parsed)
	}
	return retakes, nil
}

func (m *Model) saveSubject(idx int, code, name string) error {
	if code == "" || name == "" {
		return fmt.Errorf("code and name are required")
	}
	if idx < 0 {
		m.subjects = append(m.subjects, domain.SubjectItem{Code: code, Name: name})
		m.selectedSubj = len(m.subjects) - 1
		m.persist()
		return nil
	}
	if !inBounds(idx, len(m.subjects)) {
		return nil
	}
	m.subjects[idx].Code = code
	m.subjects[idx].Name = name
	m.persist()
	return nil
}

func (m *Model) addExam(subjectCode, examName, date, retakesRaw, priority string) error {
	if subjectCode == "" || examName == "" || date == "" {
		return fmt.Errorf("subject, exam name, and date are required")
	}
	parsedDate, err := parseStrictDateOrError(date, "date must be DD/MM/YYYY")
	if err != nil {
		return err
	}
	idx := findSubjectIndex(m.subjects, subjectCode)
	if idx < 0 {
		return fmt.Errorf("subject code not found")
	}
	retakes, err := parseRetakesOrError(retakesRaw)
	if err != nil {
		return err
	}
	m.subjects[idx].Exams = append(m.subjects[idx].Exams, domain.ExamItem{
		Name:     examName,
		Date:     parsedDate,
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
	if !inBounds(subjectIdx, len(m.subjects)) {
		return nil
	}
	exams := m.subjects[subjectIdx].Exams
	if !inBounds(examIdx, len(exams)) {
		return nil
	}
	if examName == "" || date == "" {
		return fmt.Errorf("exam name and date are required")
	}
	parsedDate, err := parseStrictDateOrError(date, "date must be DD/MM/YYYY")
	if err != nil {
		return err
	}
	retakes, err := parseRetakesOrError(retakesRaw)
	if err != nil {
		return err
	}
	exams[examIdx].Name = examName
	exams[examIdx].Date = parsedDate
	exams[examIdx].Retakes = retakes
	exams[examIdx].Priority = strings.ToUpper(priority)
	m.subjects[subjectIdx].Exams = exams
	m.examCursor = examIdx
	m.refreshExamFilter()
	m.persist()
	return nil
}

func (m *Model) saveProject(idx int, name, subject, deadline, status string) error {
	if name == "" || subject == "" || deadline == "" {
		return fmt.Errorf("name, subject, and deadline are required")
	}
	parsedDeadline, err := parseStrictDateOrError(deadline, "deadline must be DD/MM/YYYY")
	if err != nil {
		return err
	}
	if status == "" {
		status = domain.ProjectStatusNotStarted
	}
	status = strings.ToUpper(status)
	if idx < 0 {
		m.projects = append(m.projects, domain.ProjectItem{
			Name:    name,
			Subject: subject,
			Due:     parsedDeadline,
			Status:  status,
		})
		m.projectCursor = len(m.projects) - 1
		m.persist()
		return nil
	}
	if !inBounds(idx, len(m.projects)) {
		return nil
	}
	m.projects[idx].Name = name
	m.projects[idx].Subject = subject
	m.projects[idx].Due = parsedDeadline
	m.projects[idx].Status = status
	m.projectCursor = idx
	m.persist()
	return nil
}

func (m *Model) saveTodo(idx int, task, due string) error {
	if task == "" {
		return fmt.Errorf("task is required")
	}
	if due == "" {
		return fmt.Errorf("due date is required")
	}
	parsedDue, err := parseStrictDateOrError(due, "due date must be DD/MM/YYYY")
	if err != nil {
		return err
	}
	if idx < 0 {
		m.checklistItems = append(m.checklistItems, domain.ChecklistItem{
			Text: task,
			Done: false,
			Due:  parsedDue,
		})
		m.checklistCursor = len(m.checklistItems) - 1
		m.persist()
		m.refreshChecklistView()
		return nil
	}
	if !inBounds(idx, len(m.checklistItems)) {
		return nil
	}
	m.checklistItems[idx].Text = task
	m.checklistItems[idx].Due = parsedDue
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
		return m.saveTodo(-1, m.formValue(0), m.formValue(1))
	case modalEditSubject:
		return m.saveSubject(m.editSubjectIdx, m.formValue(0), m.formValue(1))
	case modalEditExam:
		return m.updateExam(m.editSubjectIdx, m.editExamIdx, m.formValue(0), m.formValue(1), m.formValue(2), m.formValue(3))
	case modalEditProject:
		return m.saveProject(m.editProjectIdx, m.formValue(0), m.formValue(1), m.formValue(2), m.formValue(3))
	case modalEditTodo:
		return m.saveTodo(m.editTodoIdx, m.formValue(0), m.formValue(1))
	case modalEditLofiURL:
		url := m.formValue(0)
		if url == "" {
			return fmt.Errorf("playlist URL is required")
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
		if !inBounds(m.checklistCursor, len(m.checklistItems)) {
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
		if inBounds(m.confirmAction.subjectIdx, len(m.subjects)) {
			m.subjects = append(m.subjects[:m.confirmAction.subjectIdx], m.subjects[m.confirmAction.subjectIdx+1:]...)
			m.selectedSubj = clampIndex(m.selectedSubj, len(m.subjects))
		}
	case confirmDeleteProject:
		if inBounds(m.confirmAction.projectIdx, len(m.projects)) {
			m.projects = append(m.projects[:m.confirmAction.projectIdx], m.projects[m.confirmAction.projectIdx+1:]...)
			m.projectCursor = clampIndex(m.projectCursor, len(m.projects))
		}
	case confirmDeleteTodo:
		if inBounds(m.confirmAction.projectIdx, len(m.checklistItems)) {
			idx := m.confirmAction.projectIdx
			m.checklistItems = append(m.checklistItems[:idx], m.checklistItems[idx+1:]...)
			m.checklistCursor = clampIndex(m.checklistCursor, len(m.checklistItems))
			m.refreshChecklistView()
		}
	case confirmClearAll:
		m.subjects = nil
		m.projects = nil
		m.checklistItems = nil
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

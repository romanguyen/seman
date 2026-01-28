package screens

import (
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	"student-exams-manager/internal/domain"
	"student-exams-manager/internal/ui/components"
)

type State struct {
	ActiveTab     int
	ConfirmOn     bool
	Checklist     viewport.Model
	Projects      []domain.ProjectItem
	Subjects      []domain.SubjectItem
	SelectedSubj  int
	ExamCursor    int
	FilteredExams []domain.ExamItem
	FocusExams    bool
	WeekLabel     string
	FilterStart   time.Time
	FilterEnd     time.Time
	FilterAll     bool
	WeekSpan      int
	WeeklyExams   []string
	ProjectCursor int
	LofiEnabled   bool
	LofiURL       string
	LofiStatus    string
	LofiError     string
	LofiPlaylist  []domain.LofiTrack
	LofiCursor    int
	LofiOffset    int
	LofiNow       int
	Modal         components.ModalState
}

package screens

import (
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	"seman/internal/domain"
	"seman/internal/ui/components"
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
	ProjectCursor int
	LofiEnabled   bool
	LofiURL       string
	LofiStatus    string
	LofiError     string
	LofiPlaylist  []domain.LofiTrack
	LofiCursor    int
	LofiNow       int
	Modal         components.ModalState
}

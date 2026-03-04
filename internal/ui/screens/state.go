package screens

import (
	"time"

	"github.com/romanguyen/seman/internal/models"
	"github.com/romanguyen/seman/internal/ui/components"
)

type State struct {
	ActiveTab     int
	ConfirmOn     bool
	ChecklistView string
	Projects      []models.ProjectItem
	Subjects      []models.SubjectItem
	SelectedSubj  int
	ExamCursor    int
	FilteredExams []models.ExamItem
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
	LofiPlaylist  []models.LofiTrack
	LofiCursor    int
	LofiOffset    int
	LofiNow       int
	Modal         components.ModalState
}

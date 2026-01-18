package screens

import (
	"time"

	"student-exams-manager/internal/models"
	"student-exams-manager/internal/ui/components"
)

type State struct {
	ActiveTab      int
	ConfirmOn      bool
	ChecklistView  string
	Projects       []models.ProjectItem
	Subjects       []models.SubjectItem
	SelectedSubj   int
	ExamCursor     int
	FilteredExams  []models.ExamItem
	FocusExams     bool
	WeekLabel      string
	FilterStart    time.Time
	FilterEnd      time.Time
	FilterAll      bool
	WeekSpan       int
	WeeklyExams    []string
	ProjectCursor  int
	Modal          components.ModalState
}

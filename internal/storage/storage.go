package storage

import "student-exams-manager/internal/models"

type SemesterData struct {
	Subjects    []models.SubjectItem   `json:"subjects"`
	Projects    []models.ProjectItem   `json:"projects"`
	Checklist   []models.ChecklistItem `json:"checklist"`
	WeeklyExams []string               `json:"weekly_exams"`
	ConfirmOn   bool                   `json:"confirm_on"`
	WeekStart   string                 `json:"week_start"`
	WeekSpan    int                    `json:"week_span"`
}

type Store interface {
	Load() (SemesterData, bool, error)
	Save(SemesterData) error
}

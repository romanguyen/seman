package storage

import "seman/internal/domain"

type SemesterData struct {
	Subjects    []domain.SubjectItem   `json:"subjects"`
	Projects    []domain.ProjectItem   `json:"projects"`
	Checklist   []domain.ChecklistItem `json:"checklist"`
	ConfirmOn   bool                   `json:"confirm_on"`
	WeekStart   string                 `json:"week_start"`
	WeekSpan    int                    `json:"week_span"`
	LofiEnabled bool                   `json:"lofi_enabled"`
	LofiURL     string                 `json:"lofi_url"`
}

type Store interface {
	Load() (SemesterData, bool, error)
	Save(SemesterData) error
}

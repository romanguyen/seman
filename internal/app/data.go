package app

import (
	"student-exams-manager/internal/domain"
	"student-exams-manager/internal/storage"
)

func DefaultData() storage.SemesterData {
	return storage.SemesterData{
		Subjects: []domain.SubjectItem{
			{
				Code: "PHY150",
				Name: "Physics",
				Exams: []domain.ExamItem{
					{
						Name:     "Test",
						Date:     "Jan 14, 2026 @ 10:00",
						Retakes:  []string{"Jan 21, 2026 @ 10:00"},
						Priority: domain.PriorityMed,
					},
					{
						Name:     "Final Exam",
						Date:     "Jan 17, 2026 @ 09:00",
						Retakes:  []string{"Jan 24, 2026 @ 09:00"},
						Priority: domain.PriorityHigh,
					},
				},
			},
			{
				Code: "HIST210",
				Name: "History",
				Exams: []domain.ExamItem{
					{
						Name:     "Oral Exam",
						Date:     "Jan 15, 2026 @ 14:00",
						Retakes:  []string{"Jan 22, 2026 @ 14:00"},
						Priority: domain.PriorityMed,
					},
				},
			},
			{
				Code: "ENG102",
				Name: "English",
				Exams: []domain.ExamItem{
					{
						Name:     "Listening Test",
						Date:     "Jan 16, 2026 @ 11:30",
						Retakes:  []string{"Jan 23, 2026 @ 11:30"},
						Priority: domain.PriorityLow,
					},
				},
			},
			{
				Code: "CS101",
				Name: "Computer Science I",
				Exams: []domain.ExamItem{
					{
						Name:     "Final Exam",
						Date:     "Jan 13, 2026 @ 13:00",
						Retakes:  []string{"Jan 20, 2026 @ 13:00"},
						Priority: domain.PriorityHigh,
					},
				},
			},
			{Code: "MATH201", Name: "Calculus II"},
		},
		Projects: []domain.ProjectItem{
			{Name: "Database Design Project", Subject: "CS101", Due: "May 20, 2025", Status: domain.ProjectStatusInProgress},
			{Name: "Renaissance Art Essay", Subject: "HIST210", Due: "May 28, 2025", Status: domain.ProjectStatusNotStarted},
			{Name: "Physics Lab Simulation", Subject: "PHY150", Due: "Jun 3, 2025", Status: domain.ProjectStatusInProgress},
			{Name: "Calculus Portfolio", Subject: "MATH201", Due: "Jun 10, 2025", Status: domain.ProjectStatusNotStarted},
			{Name: "English Literature Review", Subject: "ENG102", Due: "Jun 15, 2025", Status: domain.ProjectStatusDone},
			{Name: "Machine Learning Assignment", Subject: "CS101", Due: "Jun 18, 2025", Status: domain.ProjectStatusNotStarted},
		},
		Checklist: []domain.ChecklistItem{
			{Text: "Create flashcards for history", Done: true},
			{Text: "Attend study group session", Done: true},
			{Text: "Mock exam practice CS101", Done: false},
			{Text: "Summarize math key formulas", Done: false},
			{Text: "Review CS101 lectures 8-10", Done: false},
			{Text: "Complete MATH201 problem set 5", Done: true},
			{Text: "Read PHY150 chapter 12-13", Done: false},
			{Text: "Outline HIST210 essay", Done: false},
			{Text: "Practice ENG102 pronunciation", Done: true},
			{Text: "Work on database project", Done: false},
			{Text: "Group meeting for CS project", Done: false},
			{Text: "Submit physics lab report", Done: true},
			{Text: "Study calculus theorems", Done: false},
			{Text: "Review weak areas in physics", Done: false},
			{Text: "Prepare oral exam notes", Done: false},
		},
		WeeklyExams: []string{},
		ConfirmOn:   true,
		WeekStart:   "2026-01-12",
		WeekSpan:    1,
		LofiEnabled: false,
		LofiURL:     "",
	}
}

package components

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"seman/internal/domain"
	"seman/internal/style"
)

type upcomingExam struct {
	Subject  string
	Name     string
	Priority string
	Date     time.Time
}

func RenderUpcomingExams(subjects []domain.SubjectItem, limit int, start, end time.Time, all bool, t style.Theme) string {
	exams := collectUpcomingExams(subjects, start, end, all)
	if len(exams) == 0 {
		return t.Dim.Render("No exams this week")
	}

	if limit <= 0 || limit > len(exams) {
		limit = len(exams)
	}

	var b strings.Builder
	for i := 0; i < limit; i++ {
		if i > 0 {
			b.WriteString("\n")
		}
		item := exams[i]
		b.WriteString(t.Text.Render(fmt.Sprintf("- %s (%s)", item.Name, item.Subject)))
		b.WriteString("\n")
		dateStr := domain.FormatDate(item.Date)
		b.WriteString(t.Dim.Render("  " + dateStr + " "))
		if item.Priority != "" {
			b.WriteString(RenderPriority(item.Priority))
		}
	}

	if len(exams) > limit {
		b.WriteString("\n")
		b.WriteString(t.Dim.Render(fmt.Sprintf("...and %d more", len(exams)-limit)))
	}

	return b.String()
}

func collectUpcomingExams(subjects []domain.SubjectItem, start, end time.Time, all bool) []upcomingExam {
	list := make([]upcomingExam, 0)
	for _, subject := range subjects {
		for _, exam := range subject.Exams {
			date, ok := domain.ParseExamDate(exam.Date)
			if !ok {
				continue
			}
			if !all && (date.Before(start) || !date.Before(end)) {
				continue
			}
			list = append(list, upcomingExam{
				Subject:  subject.Code,
				Name:     exam.Name,
				Priority: exam.Priority,
				Date:     date,
			})
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Date.Before(list[j].Date)
	})
	return list
}

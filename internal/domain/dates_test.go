package domain

import "testing"

func TestParseExamDate(t *testing.T) {
	parsed, ok := ParseExamDate("Jan 2, 2006 @ 15:04")
	if !ok {
		t.Fatalf("expected exam date to parse")
	}
	if parsed.Year() != 2006 || parsed.Month().String() != "January" || parsed.Day() != 2 {
		t.Fatalf("unexpected exam date: %v", parsed)
	}

	_, ok = ParseExamDate("not-a-date")
	if ok {
		t.Fatalf("expected invalid exam date to fail")
	}
}

func TestParseTodoDate(t *testing.T) {
	parsed, ok := ParseTodoDate("2026-01-12")
	if !ok {
		t.Fatalf("expected todo date to parse")
	}
	if parsed.Year() != 2026 || parsed.Month().String() != "January" || parsed.Day() != 12 {
		t.Fatalf("unexpected todo date: %v", parsed)
	}

	_, ok = ParseTodoDate("bad")
	if ok {
		t.Fatalf("expected invalid todo date to fail")
	}
}

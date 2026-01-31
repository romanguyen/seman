package storage

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"seman/internal/domain"
)

func TestJSONStoreLoadMissing(t *testing.T) {
	dir := t.TempDir()
	store := NewJSONStore(filepath.Join(dir, "missing.json"))
	_, found, err := store.Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if found {
		t.Fatalf("expected missing file to return found=false")
	}
}

func TestJSONStoreSaveLoadRoundtrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "semester.json")
	store := NewJSONStore(path)

	data := SemesterData{
		Subjects:    []domain.SubjectItem{{Code: "CS101", Name: "Computer Science"}},
		Projects:    []domain.ProjectItem{{Name: "Final Project", Subject: "CS101", Due: "2026-05-01", Status: domain.ProjectStatusNotStarted}},
		Checklist:   []domain.ChecklistItem{{Text: "Read notes", Done: false, Due: "2026-01-12"}},
		ConfirmOn:   true,
		WeekStart:   "2026-01-12",
		WeekSpan:    2,
		LofiEnabled: false,
		LofiURL:     "",
	}

	if err := store.Save(data); err != nil {
		t.Fatalf("save failed: %v", err)
	}
	loaded, found, err := store.Load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if !found {
		t.Fatalf("expected saved file to be found")
	}
	if !reflect.DeepEqual(loaded, data) {
		payload, _ := os.ReadFile(path)
		t.Fatalf("loaded data mismatch:\nexpected: %#v\nactual: %#v\njson: %s", data, loaded, string(payload))
	}
}

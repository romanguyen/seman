package app

import (
	_ "embed"
	"encoding/json"

	"student-exams-manager/internal/storage"
)

//go:embed default_semester.json
var defaultSemesterJSON []byte

func DefaultData() storage.SemesterData {
	var data storage.SemesterData
	if err := json.Unmarshal(defaultSemesterJSON, &data); err != nil {
		return storage.SemesterData{}
	}
	return data
}

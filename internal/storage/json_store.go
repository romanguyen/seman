package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type JSONStore struct {
	path string
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{path: path}
}

func (s *JSONStore) Load() (SemesterData, bool, error) {
	file, err := os.Open(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return SemesterData{}, false, nil
		}
		return SemesterData{}, false, err
	}
	defer func() {
		_ = file.Close()
	}()

	var data SemesterData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return SemesterData{}, true, err
	}
	return data, true, nil
}

func (s *JSONStore) Save(data SemesterData) error {
	dir := filepath.Dir(s.path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	payload, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	payload = append(payload, '\n')

	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, payload, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, s.path)
}

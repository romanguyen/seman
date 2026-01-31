package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"seman/internal/app"
	"seman/internal/storage"
)

func main() {
	dataPath := flag.String("data", "", "Path to semester data JSON")
	flag.Parse()

	path, err := resolveDataPath(*dataPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving data path: %v\n", err)
		os.Exit(1)
	}

	store := storage.NewJSONStore(path)
	data, found, err := store.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading data: %v\n", err)
		os.Exit(1)
	}
	p := tea.NewProgram(app.NewModel(store, data, found), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func resolveDataPath(flagValue string) (string, error) {
	if flagValue != "" {
		return flagValue, nil
	}
	if value := os.Getenv("SEMAN_DATA"); value != "" {
		return value, nil
	}
	dataDir, err := storage.DataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dataDir, "semester.json"), nil
}

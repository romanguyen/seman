package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/app"
	"github.com/romanguyen/KEK-keep-everything-kool/internal/storage"
)

func main() {
	store := storage.NewJSONStore("data/semester.json")
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

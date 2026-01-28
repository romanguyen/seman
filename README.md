Student Exams Manager (Go TUI)

Requirements:
- Go 1.21+
- Optional for Lofi tab: mpv, yt-dlp

Run:
- go run ./cmd/student-exams-manager
- go run ./cmd/student-exams-manager -data /path/to/semester.json
- STUDENT_EXAMS_DATA=/path/to/semester.json go run ./cmd/student-exams-manager

Data file:
- Default directory: $XDG_DATA_HOME/student-exams-manager
- Fallback directory: $HOME/.local/share/student-exams-manager
- File name: semester.json
- Auto-created on first save

Tests:
- go test ./...

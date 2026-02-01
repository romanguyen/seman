Seman (Go TUI)

A terminal UI to track subjects/exams, projects, and todos with a weekly view. Includes an optional Lofi player tab.

Requirements

- Go 1.21+
- Optional (Lofi tab): mpv and yt-dlp

Install

- Run without installing:
  - go run ./cmd/seman
- Build a local binary:
  - go build -o seman ./cmd/seman
- Install to ~/.local/bin:
  - ./install.sh
  - Note: add ~/.local/bin to your PATH if needed.

Usage

- Default:
  - seman
- Custom data path:
  - seman -data /path/to/semester.json
  - SEMAN_DATA=/path/to/semester.json seman
  - Flag takes precedence over env var.

Data storage

- Default directory: $XDG_DATA_HOME/seman
- Fallback directory: $HOME/.local/share/seman
- File name: semester.json
- The file is created on first save.
- The app starts with empty data when no file exists.

Key bindings (global)

- [1-6] switch tabs
- [A] add (contextual: Exams/Todos/Projects)
- [E] edit current item
- [D] delete current item
- [Q] quit
- [Left]/[Right] shift week

Settings tab

- [O] toggle confirm on delete
- [W] cycle week span
- [L] toggle Lofi tab
- [U] edit Lofi playlist URL
- [C] clear all data

Lofi tab

- [Enter] play current track
- [Space] play/pause
- [N] next
- [B] previous
- [x] stop

Troubleshooting (Lofi)

- If playback sticks on "Loading...", update yt-dlp and ensure mpv uses it.
- Example: mpv --no-video --script-opts=ytdl_hook-ytdl_path=yt-dlp <playlist_url>

Uninstall

- ./uninstall.sh
- ./uninstall.sh --data (removes data directory)

Tests

- go test ./...

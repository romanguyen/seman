# Seman

A terminal UI to track subjects/exams, projects, and todos with a weekly view. Includes an optional Lofi player tab.

## Requirements
- Go 1.21+
- Optional (Lofi tab): `mpv`, `yt-dlp`

## Install
Run without installing:
```bash
go run ./cmd/seman
```

Build a local binary:
```bash
go build -o seman ./cmd/seman
```

Install to `~/.local/bin`:
```bash
./install.sh
```
If `~/.local/bin` is not in your PATH, add it:
```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

## Usage
Default:
```bash
seman
```

Custom data path:
```bash
seman -data /path/to/semester.json
SEMAN_DATA=/path/to/semester.json seman
```
The `-data` flag takes precedence over `SEMAN_DATA`.

## Data Storage
- Default directory: `$XDG_DATA_HOME/seman`
- Fallback directory: `$HOME/.local/share/seman`
- File name: `semester.json`
- The file is created on first save.
- The app starts with empty data when no file exists.

## Key Bindings
### Global
| Key | Action |
| --- | ------ |
| `1-6` | switch tabs |
| `A` | add (contextual: Exams/Todos/Projects) |
| `S` | add subject (Exams tab only) |
| `E` | edit current item |
| `D` | delete current item |
| `Q` | quit |
| `Left/Right` | shift week |
| `Tab` | toggle focus (Subjects/Exams on Exams tab) |

### Settings Tab
| Key | Action |
| --- | ------ |
| `O` | toggle confirm on delete |
| `W` | cycle week span |
| `L` | toggle Lofi tab |
| `U` | edit Lofi playlist URL |
| `C` | clear all data |

### Lofi Tab
| Key | Action |
| --- | ------ |
| `Enter` | play current track |
| `Space` | play/pause |
| `N` | next |
| `B` | previous |
| `X` | stop |

## Date Format
All dates are **DD/MM/YYYY**.

## Troubleshooting (Lofi)
If playback sticks on "Loading...", update `yt-dlp` and ensure `mpv` uses it:
```bash
mpv --no-video --script-opts=ytdl_hook-ytdl_path=yt-dlp <playlist_url>
```

## Uninstall
```bash
./uninstall.sh
./uninstall.sh --data
```

## Tests
```bash
go test ./...
```

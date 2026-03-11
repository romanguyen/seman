package components

import (
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/muesli/reflow/ansi"
)

// PlaceOverlay renders fg centered on top of bg without replacing the background.
func PlaceOverlay(bg, fg string) string {
	bgLines := strings.Split(bg, "\n")
	fgLines := strings.Split(fg, "\n")

	bgH := len(bgLines)
	fgH := len(fgLines)
	if fgH == 0 {
		return bg
	}

	fgW := 0
	for _, l := range fgLines {
		if w := ansi.PrintableRuneWidth(l); w > fgW {
			fgW = w
		}
	}
	bgW := 0
	for _, l := range bgLines {
		if w := ansi.PrintableRuneWidth(l); w > bgW {
			bgW = w
		}
	}

	x := (bgW - fgW) / 2
	y := (bgH - fgH) / 2
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}

	result := make([]string, bgH)
	copy(result, bgLines)
	for i, fgLine := range fgLines {
		bgIdx := y + i
		if bgIdx >= bgH {
			break
		}
		result[bgIdx] = overlayLine(result[bgIdx], fgLine, x)
	}
	return strings.Join(result, "\n")
}

// overlayLine paints fgLine on top of bgLine starting at visible column x.
func overlayLine(bg, fg string, x int) string {
	left := ansiTruncate(bg, x)
	// pad if bg was shorter than x
	if pad := x - ansi.PrintableRuneWidth(left); pad > 0 {
		left += strings.Repeat(" ", pad)
	}
	fgW := ansi.PrintableRuneWidth(fg)
	right := ansiSkip(bg, x+fgW)
	return left + fg + right
}

// ansiTruncate returns the first n visible columns of s, preserving ANSI codes.
func ansiTruncate(s string, n int) string {
	if n <= 0 {
		return ""
	}
	var buf strings.Builder
	visible := 0
	runes := []rune(s)
	for i := 0; i < len(runes); {
		if seq, seqLen := scanAnsiSeq(runes, i); seqLen > 0 {
			buf.WriteString(seq)
			i += seqLen
			continue
		}
		r := runes[i]
		w := runewidth.RuneWidth(r)
		if visible+w > n {
			break
		}
		buf.WriteRune(r)
		visible += w
		i++
	}
	return buf.String()
}

// ansiSkip returns s with the first n visible columns removed, preserving ANSI codes.
func ansiSkip(s string, n int) string {
	var buf strings.Builder
	visible := 0
	runes := []rune(s)
	for i := 0; i < len(runes); {
		if seq, seqLen := scanAnsiSeq(runes, i); seqLen > 0 {
			if visible >= n {
				buf.WriteString(seq)
			}
			i += seqLen
			continue
		}
		r := runes[i]
		w := runewidth.RuneWidth(r)
		if visible >= n {
			buf.WriteRune(r)
		}
		visible += w
		i++
	}
	return buf.String()
}

// scanAnsiSeq returns the escape sequence starting at runes[i] and its rune length,
// or ("", 0) if runes[i] is not the start of an escape sequence.
func scanAnsiSeq(runes []rune, i int) (string, int) {
	if runes[i] != '\x1b' || i+1 >= len(runes) {
		return "", 0
	}
	// CSI sequence: ESC '[' ... <final byte 0x40-0x7E>
	if runes[i+1] == '[' {
		j := i + 2
		for j < len(runes) {
			r := runes[j]
			j++
			if r >= 0x40 && r <= 0x7E {
				return string(runes[i:j]), j - i
			}
		}
		return string(runes[i:j]), j - i
	}
	// Any other ESC sequence: consume ESC + next rune
	return string(runes[i : i+2]), 2
}

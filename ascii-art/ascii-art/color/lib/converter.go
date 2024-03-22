package ascii_art

import (
	"strings"
)

func AppendChar(a [8]string, b [8]string) [8]string {
	c := [8]string{}
	for i := 0; i < 8; i++ {
		c[i] = a[i] + b[i]
	}
	return c
}

func IsLineEmpty(line [8]string) bool {
	return line == [8]string{}
}

func StylizeColored(s string, artset []string) [][8]string {
	out := [][8]string{}
	line := [8]string{}

	// Replace the literal "\n" into newline escapes chars
	s = strings.ReplaceAll(s, "\\n", "\n")
	s = strings.ReplaceAll(s, string([]byte{0x0D, 0x0A}), "\n")

	// For ommiting unicodes characters without ascii-art-ing the bytes of the rune
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {

		if (i + 1) == 3 {
			line = AppendChar(line, [8]string{
				"\u001b[31m", "\u001b[31m", "\u001b[31m", "\u001b[31m",
				"\u001b[31m", "\u001b[31m", "\u001b[31m", "\u001b[31m",
			})
		} else if (i + 1) == 4 {
			line = AppendChar(line, [8]string{
				"\u001b[0m", "\u001b[0m", "\u001b[0m", "\u001b[0m",
				"\u001b[0m", "\u001b[0m", "\u001b[0m", "\u001b[0m",
			})
		}

		if s[i] == '\n' {
			// Store the ascii-art line when a newline is encountered
			out = append(out, line)
			// Clear the current line buffer
			line = [8]string{}
			// Skip the newline
			continue
		}
		// Get the ascii-art of this character with the artset
		a := GetArtistic(runes[i], artset)
		// Append this character
		line = AppendChar(line, a)
	}

	// Append the line even if it doesn't end with a new line
	if !IsLineEmpty(line) {
		out = append(out, line)
	}

	return out
}

package ascii_art

import (
	"fmt"
	"regexp"
	"strings"
	"syscall"
	"unsafe"
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

func Stylize(s string, artset []string) [][8]string {
	// description	: parse a string and convert that string into ascii-art with
	//				  the artset specified
	// param s		: string to convert to ascii-art
	// param artset	: artset to use (previously parsed via `ParseArtSet`)
	// return		: a list of words, where words are a list (height limit 8) of strings

	out := [][8]string{}
	line := [8]string{}

	// Replace the literal "\n" into newline escapes chars
	s = strings.ReplaceAll(s, "\\n", "\n")
	s = strings.ReplaceAll(s, string([]byte{0x0D, 0x0A}), "\n")

	// For ommiting unicodes characters without ascii-art-ing the bytes of the rune
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		// Support for newline characters

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

func GeneratePadding(padding int) string {
	out := ""
	for i := 0; i < padding; i++ {
		out += " "
	}
	return out
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func TerminalWidth() int {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return int(ws.Col)
}

// implement newlines

func ConvertRune(c rune, banner []string) [8]string {
	if c < 32 || c >= 127 {
		fmt.Println("illegal non-printable/unicode character in input")
		return [8]string{}
	}

	char := [8]string{}
	start := (c - 32) * 8
	for i, ri := start, 0; i < start+8; i, ri = i+1, ri+1 {
		char[ri] = banner[i]
	}
	return char
}

func CharacterWidth(ch [8]string) int {
	return len(ch[0])
}

func ConvertLine(content string, padding int, banner []string) [8]string {
	characters := [8]string{}

	for _, char := range content {
		if char == ' ' {
			for i := 0; i < padding; i++ {
				characters = AppendChar(characters, [8]string{
					" ", " ", " ", " ", " ", " ", " ", " ",
				})
			}
			continue
		}

		converted := ConvertRune(char, banner)
		characters = AppendChar(characters, converted)
	}

	return characters
}

func StylizeJustify(content string, artset []string) [][8]string {
	// Replace the literal "\n" into newline escapes chars
	out := [][8]string{}

	content = strings.ReplaceAll(content, string([]byte{0x0D, 0x0A}), "\\n")
	lines := strings.Split(content, "\\n")

	succession := regexp.MustCompile(" +")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = succession.ReplaceAllStringFunc(line, func(s string) string {
			return " "
		})

		spaces := 0
		non_spaces := 0

		for _, char := range line {
			if char == ' ' {
				spaces++
			} else {
				giant := ConvertRune(char, artset)
				non_spaces += CharacterWidth(giant)
			}
		}

		line_length := non_spaces + spaces
		released := TerminalWidth() - line_length

		remaining := released % spaces
		padding := released / spaces

		_ = remaining

		converted := ConvertLine(line, padding, artset)
		out = append(out, converted)
	}

	return out
}

package ascii_art

import (
	"bufio"
	"os"
)

func ParseArtSet(filepath string) []string {
	set := []string{}

	file, err := os.Open(filepath)
	if err != nil {
		panic("file not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			set = append(set, line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic("error scanning the artset")
	}

	return set
}

func GetArtistic(c rune, artset []string) [8]string {
	if c < 32 || c >= 127 {
		panic("illegal non-printable characters in input")
	}

	char := [8]string{}
	start := (c - 32) * 8
	for i, ri := start, 0; i < start+8; i, ri = i+1, ri+1 {
		char[ri] = artset[i]
	}

	return char
}

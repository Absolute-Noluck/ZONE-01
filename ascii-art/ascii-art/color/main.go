package main

import (
	"flag"
	"fmt"
	"os"

	processor "ascii-art/lib"
)

func usage() {
	fmt.Println("Usage: go run . [STRING] [OPTION]")
	fmt.Println("")
	fmt.Println("EX: go run . something --color=<color>")
}

func main() {
	if len(os.Args) != 3 {
		usage()
		return
	}

	color := flag.String("color", "", "set the color (enclosed by two '~' characters)")
	flag.CommandLine.Parse(os.Args[2:])

	args := os.Args[1:]
	text := args[0]

	if text == "" || *color == "" {
		usage()
		return
	}

	artset := processor.ParseArtSet("../banners/standard.txt")

	contentLines := processor.StylizeColored(text, artset)
	for _, line := range contentLines {
		for _, charLine := range line {
			fmt.Println(charLine)
		}
	}
}

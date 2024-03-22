package main

import (
	"fmt"
	"os"

	processor "ascii-art/lib"
)

func usage() {
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println("")
	fmt.Println("EX: go run . something standard")
}

func main() {
	if len(os.Args) != 3 {
		usage()
		return
	}

	args := os.Args[1:]
	text, bannerType := args[0], args[1]

	if text == "" || bannerType == "" {
		usage()
		return
	}

	artset := []string{}

	switch bannerType {
	case "standard":
		{
			artset = processor.ParseArtSet("../banners/standard.txt")
		}
	case "shadow":
		{
			artset = processor.ParseArtSet("../banners/shadow.txt")
		}
	case "thinkertoy":
		{
			artset = processor.ParseArtSet("../banners/thinkertoy.txt")
		}
	}

	contentLines := processor.Stylize(text, artset)
	for _, line := range contentLines {
		for _, charLine := range line {
			fmt.Println(charLine)
		}
	}
}

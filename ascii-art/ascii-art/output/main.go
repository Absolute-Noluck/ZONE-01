package main

import (
	"flag"
	"fmt"
	"os"

	processor "ascii-art/lib"
)

func usage() {
	fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
	fmt.Println("")
	fmt.Println("EX: go run . something standard --output=<fileName.txt>")
}

func main() {
	if len(os.Args) != 4 {
		usage()
		return
	}

	output := flag.String("output", "", "set the output file")
	flag.CommandLine.Parse(os.Args[3:])

	args := os.Args[1:]
	text, bannerType := args[0], args[1]

	if text == "" || bannerType == "" || *output == "" {
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
	content := ""
	for _, line := range contentLines {
		for _, charLine := range line {
			content += charLine + "\n"
		}
	}

	if os.WriteFile(*output, []byte(content), 0o600) != nil {
		fmt.Printf("cannot write file: '%s'\n", *output)
	}
}

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
	fmt.Println("EX: go run . something standard --align=right")
}

// Issue:		This is a terminal. A terminal doesn't support floating point positions.
//		 		How to equally space words with the same amount of spaces in these conditions ?
//        	 	Adding a space to reach the margin-right will just fuck up the justification and
//        	 	the fixed-size padding.
// Solution: 	Use normal spaces unstead of the ascii-art ones.

// Issue:		Some spaces remains, ones that can't be filled without putting the spaces
//				(number of spaces) in disarray.
// Solution:    Do not add the remaining spaces anyway, the user will need to shrink
//				their terminal's width. Only after that the text will be justified to both
//				borders of the terminal.

func main() {
	if len(os.Args) != 4 {
		usage()
		return
	}

	alignment := flag.String("align", "", "set the alignment file")
	flag.CommandLine.Parse(os.Args[3:])

	args := os.Args[1:]
	text, bannerType := args[0], args[1]

	if text == "" || bannerType == "" || *alignment == "" {
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
	default:
		{
			fmt.Println("invalid banner: not found")
		}
	}

	if !(*alignment == "left" || *alignment == "right" || *alignment == "center" || *alignment == "justify") {
		fmt.Println("invalid alignment; must be left, right, center or justify")
		return
	}

	var contentLines [][8]string
	if *alignment == "justify" {
		contentLines = processor.StylizeJustify(text, artset)
	} else {
		contentLines = processor.Stylize(text, artset)
	}

	for _, line := range contentLines {
		for _, charLine := range line {

			padding := 0
			switch *alignment {
			case "right":
				padding = processor.TerminalWidth() - len(charLine) - 1
			case "center":
				padding = (processor.TerminalWidth() / 2) - (len(charLine) / 2)
			}

			fmt.Println(processor.GeneratePadding(padding), charLine)
		}
	}
}

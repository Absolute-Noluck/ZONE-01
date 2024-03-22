package main

import (
	"fmt"
	"os"

	processor "ascii-art/lib"
)

const ARTSET_USE = "../banners/standard.txt"

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		arg := args[0]

		artset := processor.ParseArtSet(ARTSET_USE)

		output := processor.Stylize(arg, artset)
		for _, line := range output {
			for _, subline := range line {
				fmt.Println(subline)
			}
		}
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ascii(sentence, banner string) {
	// args := os.Args[1]
	var charLine string
	// var sentence2 string
	// if len(args) == 2 {
	os.Remove("test.txt")
	file2, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0o600)
	defer file2.Close() // on ferme automatiquement à la fin de notre programme
	sentence = strings.ReplaceAll(sentence, "\r", "\n")
	line := strings.Split(sentence, "\n")
	/*for _, ch := range line {
		fmt.Println(ch)
		sentence2 += ch
	}*/
	//line2 := strings.Split(sentence2, "\r")
	//fmt.Println(len(line))
	// fmt.Println(len(line2))
	content := getTable(banner)
	// for y := 0; y < len(line2); y++ {
	for i := 0; i < len(line); i++ {
		if len(line[i]) > 0 {
			chars := []rune(line[i])
			// fmt.Println(len(chars))
			for n := 0; n < 8; n++ {
				for v := 0; v < len(chars); v++ {
					char := (int(chars[v]) - 32) * 9
					charLine = content[char+1+n]
					_, err := file2.WriteString(string(charLine)) // écrire dans le fichier
					if err != nil {
						panic(err)
					}
				}
				_, err := file2.WriteString("\n") // écrire dans le fichier
				if err != nil {
					panic(err)
				}
			}
		}
	}
	//}
	/*} else if len(args) == 3 {
		line := strings.Split(os.Args[1], "\\n")
		content := getTable()
		for i := 0; i < len(line); i++ {
			if len(line[i]) > 0 {
				chars := []rune(line[i])
				for n := 0; n < 8; n++ {
					for v := 0; v < len(chars); v++ {
						char := (int(chars[v]) - 32) * 9
						charLine := content[char+1+n]
						fmt.Print(charLine)
					}
					fmt.Print(string(rune('\n')))
				}
			} else {
				fmt.Print(string(rune('\n')))
			}
		}
	}*/
}

func getTable(banner string) []string {
	var table []string
	/*file, err := os.Open("standard.txt") // lire le fichier text.txt
	content, _ := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	table = strings.Split(string(content), "\n")*/
	// args := os.Args[1:]
	// if len(args) == 1 {
	if banner == "standard" {
		file, err := os.Open("template/standard.txt") // lire le fichier text.txt
		content, _ := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		table = strings.Split(string(content), "\n")
	} else if banner == "thinkertoy" {
		file, err := os.Open("template/Thinkertoy.txt") // lire le fichier text.txt
		content, _ := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		table = strings.Split(string(content), "\n")
	} else if banner == "shadow" {
		file, err := os.Open("template/shadow.txt") // lire le fichier text.txt
		content, _ := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		table = strings.Split(string(content), "\n")
	}
	/*} else if len(args) == 2 {
		if string(args[1]) == "standard" {
			file, err := os.Open("standard.txt") // lire le fichier text.txt
			content, _ := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			table = strings.Split(string(content), "\n")
		}
		if string(args[1]) == "Shadows" {
			file, err := os.Open("Shadows.txt") // lire le fichier text.txt
			content, _ := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			table = strings.Split(string(content), "\n")
		}
		if string(args[1]) == "Tinker" {
			file, err := os.Open("Tinker.txt") // lire le fichier text.txt
			content, _ := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			table = strings.Split(string(content), "\n")
		}
	}*/
	return table
}

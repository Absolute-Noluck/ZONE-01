package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// open file
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read the file word by word using scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var s []string
	for scanner.Scan() {
		s = append(s, scanner.Text()) // récupération en Tableau
	}
	var sortie string
	// Conditions
	var count int
	for i := 0; i < len(s)-1; i++ {
		for a := 0; a < len(s)-1; a++ {
			if s[i] == "a" || s[i] == "A" {
				word := s[i+1]
				if string(word[0]) == "a" || string(word[0]) == "A" || string(word[0]) == "e" || string(word[0]) == "E" || string(word[0]) == "H" || string(word[0]) == "h" || string(word[0]) == "I" || string(word[0]) == "i" || string(word[0]) == "O" || string(word[0]) == "o" || string(word[0]) == "U" || string(word[0]) == "u" || string(word[0]) == "Y" || string(word[0]) == "y" {
					s[i] = s[i] + "n"
				}
			}

			if s[a] == "(cap)" {
				h := cap(s[count-1])
				s[count] = h
				s = RemoveIndex(s, count-1)
			} else if s[a] == "(cap," {
				b := strings.Trim(string(s[count+1]), s[count+1][1:])
				number, _ := strconv.Atoi(string(b))
				for j := 1; j <= number; j++ {
					s[count-j] = cap(s[count-j])
				}
				s = append(s[:count], s[count+2:]...)
			}

			if s[a] == "(up)" {
				h := ToUpper(s[count-1])
				s[count] = h
				s = RemoveIndex(s, count-1)
			} else if s[a] == "(up," {
				b := strings.Trim(string(s[count+1]), s[count+1][1:])
				number, _ := strconv.Atoi(string(b))
				for j := 1; j <= number; j++ {
					s[count-j] = ToUpper(s[count-j])
				}
				s = append(s[:count], s[count+2:]...)
			}

			if s[a] == "(low)" {
				h := ToLower(s[count-1])
				s[count] = h
				s = RemoveIndex(s, count-1)
			} else if s[a] == "(low," {
				b := strings.Trim(string(s[count+1]), s[count+1][1:])
				number, _ := strconv.Atoi(string(b))
				for j := 1; j <= number; j++ {
					s[count-j] = ToLower(s[count-j])
				}
				s = append(s[:count], s[count+2:]...)
			}
			if s[a] == "(hex)" {
				h := Todec(s[count-1])
				s[count] = h
				s = RemoveIndex(s, count-1)
			} else if s[a] == "(hex," {
				b := strings.Trim(string(s[count+1]), s[count+1][1:])
				number, _ := strconv.Atoi(string(b))
				for j := 1; j <= number; j++ {
					s[count-j] = Todec(s[count-j])
				}
				s = append(s[:count], s[count+2:]...)
			}

			if s[a] == "(bin)" {
				h := Bin(s[count-1])
				s[count] = h
				s = RemoveIndex(s, count-1)
			} else if s[a] == "(bin," {
				b := strings.Trim(string(s[count+1]), s[count+1][1:])
				number, _ := strconv.Atoi(string(b))
				for j := 1; j <= number; j++ {
					s[count-j] = Bin(s[count-j])
				}
				s = append(s[:count], s[count+2:]...)
			}

			count++
		}
	}
	for a, b := range s {
		if b != "," && b != "." && b != "?" && b != "!" && b != ":" && b != ";" && b != "..." && b != "!?" && a > 0 {
			sortie += " "
		}
		sortie += b

	}
	ExpressionA := regexp.MustCompile(`([']([^']+)['])`)
	sortie = ExpressionA.ReplaceAllStringFunc(sortie, func(s string) string {
		rs := []rune(s)
		ss := rs[1 : len(rs)-1]
		return "'" + strings.TrimSpace(string(ss)) + "'"
	})

	ExpressionAI := regexp.MustCompile(`([‘]([^‘]+)[‘])`)
	sortie = ExpressionAI.ReplaceAllStringFunc(sortie, func(s string) string {
		rs := []rune(s)
		ss := rs[1 : len(rs)-1]
		return "‘" + strings.TrimSpace(string(ss)) + "‘"
	})

	ExpressionV := regexp.MustCompile(`(?:\s+)?,(?:\s+)?`)
	sortie = ExpressionV.ReplaceAllStringFunc(sortie, func(v string) string {
		return ", "
	})
	f, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(sortie)
	if err != nil {
		fmt.Println(err)
		f.Close()
		l = l
		return
	} else {
		fmt.Println("Your file is created !")
	}
}

// CAP
func cap(s string) string {
	element := s
	var result []byte
	for _, a := range element {
		result = append(result, byte(a))
	}
	result[0] = result[0] - 32
	final := string(result)
	return final
}

// LOW
func ToLower(s string) string {
	chaine := []rune(s)
	for i := 0; i < len(chaine); i++ {
		c := chaine[i]
		if 'A' <= c && c <= 'Z' {
			c -= 'A' - 'a'
		}
		chaine[i] = c
	}
	return string(chaine)
}

// UP
func ToUpper(s string) string {
	chaine := []byte(s)
	for i := 0; i < len(chaine); i++ {
		c := chaine[i]
		if 97 <= c && c <= 122 {
			c -= 97 - 65
		}
		chaine[i] = c
	}
	return string(chaine)
}

// HEX
func Todec(s string) string {
	nombre, _ := strconv.ParseInt(s, 16, 64)
	return fmt.Sprint(nombre)
}

func Bin(s string) string {
	nombre, _ := strconv.ParseInt(s, 2, 64)
	return fmt.Sprint(nombre)
}

//
// Clean
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

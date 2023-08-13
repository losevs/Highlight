package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/fatih/color"
)

func countLetters(str string) int {
	count := 0
	for _, letter := range str {
		if unicode.IsLetter(letter) {
			count++
		}
	}
	return count
}

func main() {
	file, err := os.Open("Lorem.txt")
	if err != nil {
		log.Fatalln("cannot open file", err)
	}
	defer file.Close()

	text := ""

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		for ind := range scanner.Text() {
			switch {
			case countLetters(scanner.Text()) < 4:
				if ind == 0 {
					fmt.Print(color.BlackString(string(scanner.Text()[ind])))

					text += color.BlackString(string(scanner.Text()[ind]))
				} else {
					fmt.Print(color.HiBlackString(string(scanner.Text()[ind])))

					text += string(scanner.Text()[ind])
				}

			case countLetters(scanner.Text()) == 4:
				if ind == 0 || ind == 1 {
					fmt.Print(color.BlackString(string(scanner.Text()[ind])))

					text += color.BlackString(string(scanner.Text()[ind]))

				} else {
					fmt.Print(color.HiBlackString(string(scanner.Text()[ind])))

					text += string(scanner.Text()[ind])
				}

			case countLetters(scanner.Text()) <= 6:
				if ind == 0 || ind == 1 || ind == 2 {
					fmt.Print(color.BlackString(string(scanner.Text()[ind])))

					text += color.BlackString(string(scanner.Text()[ind]))
				} else {
					fmt.Print(color.HiBlackString(string(scanner.Text()[ind])))

					text += string(scanner.Text()[ind])
				}

			default:
				if ind == 0 || ind == 1 || ind == 2 || ind == 3 {
					fmt.Print(color.BlackString(string(scanner.Text()[ind])))

					text += color.BlackString(string(scanner.Text()[ind]))
				} else {
					fmt.Print(color.HiBlackString(string(scanner.Text()[ind])))

					text += string(scanner.Text()[ind])
				}
			}

		}
		if strings.HasSuffix(scanner.Text(), ".") {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
		text += " "
	}
}

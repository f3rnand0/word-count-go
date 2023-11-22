package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	var args = os.Args
	var text, result, command, file string

	if len(args) == 1 {
		text = getTextFromInput()
		if len(text) == 0 {
			fmt.Println("Usage: application [-c] filename")
			os.Exit(1)
		} else {
			result = countLines(text) + " " + countWords(text) + " " + countCharacters(text)
			fmt.Printf("%s\n", result)
		}
	} else {
		text = getTextFromInput()

		if len(text) == 0 {
			command = args[1]
			if !strings.Contains(command, "-") {
				file = command
				command = ""
			} else {
				file = args[2]
			}

			text = getTextFromFile(file)

			switch command {
			case "-c":
				result = countBytes(text)
			case "-l":
				result = countLines(text)
			case "-w":
				result = countWords(text)
			case "-m":
				result = countCharacters(text)
			default:
				result = countLines(text) + " " + countWords(text) + " " + countCharacters(text)
			}
			fmt.Printf("%s %s\n", result, file)

		} else {
			command = args[1]

			switch command {
			case "-c":
				result = countBytes(text)
			case "-l":
				result = countLines(text)
			case "-w":
				result = countWords(text)
			case "-m":
				result = countCharacters(text)
			}
			fmt.Printf("%s\n", result)
		}
	}

}

func countBytes(text string) string {
	return strconv.Itoa(utf8.RuneCountInString(text))
}

func countLines(text string) string {
	counter := 0
	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		counter++
	}
	return strconv.Itoa(counter)
}

func countWords(text string) string {
	counter := 0
	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words := len(strings.Fields(line))
		counter += words
	}
	return strconv.Itoa(counter)
}

func countCharacters(text string) string {
	counter := 0
	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		chars := utf8.RuneCountInString(line)
		counter += chars
	}
	return strconv.Itoa(counter)
}

func getTextFromInput() string {
	fi, err := os.Stdin.Stat()
	if err != nil {
		os.Exit(1)
	}
	size := fi.Size()
	if size > 0 {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			os.Exit(1)
		}
		return string(bytes[:])
	} else {
		return ""
	}

}

func getTextFromFile(file string) string {
	fi, err := os.ReadFile(file)
	if err != nil {
		os.Exit(1)
	}

	return string(fi)
}

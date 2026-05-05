package main

import (
	"os"
	"strings"
	"fmt"
)

func GenerateAscii(text, banner string) string {
	
	data, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		return "Error reading banner file\n"
	}

	fileLines := strings.Split(string(data), "\n")
	result := ""
	words := strings.Split(text, "\n")

	for _, word := range words {
		for i := 1; i <= 8; i++ {
			for _, char := range word {
				asciiIndex := int(char) - 32
				start := asciiIndex * 9

				if start+i >= len(fileLines) {
					continue
				}

				result += fileLines[start+i]
			}
			result += "\n"
		}
	}
	return result
}

func PrintJustified(words []string, banner string, width int) {

	if len(words) == 0 {
		return
	}

	// build ASCII for each word separately
	asciiWords := make([][]string, len(words))

	maxHeight := 0

	for i, w := range words {
		ascii := strings.Split(GenerateAscii(w, banner), "\n")
		asciiWords[i] = ascii
		if len(ascii) > maxHeight {
			maxHeight = len(ascii)
		}
	}

	// print line by line (row of ASCII art)
	for row := 0; row < maxHeight; row++ {

		var lineParts []string

		for i := 0; i < len(asciiWords); i++ {
			if row < len(asciiWords[i]) {
				lineParts = append(lineParts, asciiWords[i][row])
			}
		}

		if len(lineParts) == 0 {
			fmt.Println()
			continue
		}

		// compute total length
		totalLen := 0
		for _, p := range lineParts {
			totalLen += len(p)
		}

		gaps := len(lineParts) - 1

		if gaps <= 0 {
			fmt.Println(lineParts[0])
			continue
		}

		spaces := width - totalLen
		if spaces < gaps {
			spaces = gaps
		}

		baseSpace := spaces / gaps
		extra := spaces % gaps

		var result string

		for i, part := range lineParts {
			result += part

			if i < gaps {
				result += strings.Repeat(" ", baseSpace)

				if extra > 0 {
					result += " "
					extra--
				}
			}
		}

		fmt.Println(result)
	}
}
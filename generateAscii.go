package main

import (
	"os"
	"strings"
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

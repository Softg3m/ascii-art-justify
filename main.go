package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--align" {
		Usage()
		return
	}

	// WITH ALIGN FLAG
	if strings.HasPrefix(args[0], "--align=") {
		alignType := strings.TrimPrefix(args[0], "--align=")

		// validate align type
		if alignType != "left" && alignType != "right" && alignType != "center" && alignType != "justify" {
			Usage()
			return
		}

		// check required args
		if len(args) < 2 {
			Usage()
			return
		}

		text := args[1]
		banner := "standard"

		if len(args) == 3 {
			banner = args[2]
		}

		ascii := GenerateAscii(text, banner)
		lines := strings.Split(ascii, "\n")

		width := 150

		for _, line := range lines {
			switch alignType {

			case "left":
				fmt.Println(line)

			case "right":
				spaces := width - len(line)
				if spaces > 0 {
					fmt.Println(strings.Repeat(" ", spaces) + line)
				} else {
					fmt.Println(line)
				}

			case "center":
				spaces := (width - len(line)) / 2
				if spaces > 0 {
					fmt.Println(strings.Repeat(" ", spaces) + line)
				} else {
					fmt.Println(line)
				}

			case "justify":
				fmt.Println(line)
			}
		}

		// WITHOUT ALIGN FLAG
	} else {
		text := args[0]
		banner := "standard"

		if len(args) == 2 {
			banner = args[1]
		}

		ascii := GenerateAscii(text, banner)
		fmt.Print(ascii)
	}
}

func Usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("Example: go run . --align=right something standard")
}

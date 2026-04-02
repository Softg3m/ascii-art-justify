package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		Usage()
		return
	}

	if strings.HasPrefix(args[0], "--align=") {
		alignType := strings.TrimPrefix(args[0], "--align=")

		if alignType != "left" && alignType != "right" && alignType != "center" && alignType != "justify" {
			Usage()
			return
		}

		text := args[1]
		banner := "standard"

		if len(args) == 3 {
			banner = args[2]
		}
		ascii := GenerateAscii(text, banner)
		fmt.Print(ascii)

	} else {
		text := args[0]
		banner := "standard"

		if len(args) == 2 {
			banner = args[1]
		}
		ascii := GenerateAscii(text, banner)
		fmt.Println(ascii)
	}

}

func Usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("Example: go run . --align=right something standard")
}

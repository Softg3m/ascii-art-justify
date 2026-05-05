package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		Usage()
		return
	}

	alignType := "left"
	textIndex := 0
	banner := "standard"

	// handle optional flag
	if strings.HasPrefix(args[0], "--align=") {
		alignType = strings.TrimPrefix(args[0], "--align=")

		if alignType != "left" && alignType != "right" && alignType != "center" && alignType != "justify" {
			Usage()
			return
		}

		textIndex = 1
	}

	if len(args) <= textIndex {
		Usage()
		return
	}

	text := args[textIndex]
	text = strings.ReplaceAll(text, "\\n", "\n")

	if len(args) > textIndex+1 {
		banner = args[textIndex+1]
	}

	// JUSTIFY MUST WORK ON WORD LEVEL
	words := strings.Fields(text)

	width := getTerminalWidth()

	if alignType == "justify" {
		PrintJustified(words, banner, width)
		return
	}

	ascii := GenerateAscii(text, banner)
	lines := strings.Split(ascii, "\n")

	for _, line := range lines {
		if line == "" {
			return
		}

		switch alignType {

		case "left":
			fmt.Println(line)

		case "right":
			padding := width - len(line)
			if padding > 0 {
				fmt.Println(strings.Repeat(" ", padding) + line)
			} else {
				fmt.Println(line)
			}

		case "center":
			padding := (width - len(line)) / 2
			if padding > 0 {
				fmt.Println(strings.Repeat(" ", padding) + line)
			} else {
				fmt.Println(line)
			}
		}
	}
}

func Usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("Example: go run . --align=right something standard")
}

func getTerminalWidth() int {
	type winsize struct {
		Row, Col, Xpixel, Ypixel uint16
	}

	ws := &winsize{}
	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(0),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if err != 0 {
		return 80
	}
	return int(ws.Col)
}

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

		width := getTerminalWidth()

		for _, line := range lines {
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

func getTerminalWidth() int {
	type winsize struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}

	ws := &winsize{}
	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(0),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)
	if err != 0{
		return 80
	}

	return int(ws.Col)
}

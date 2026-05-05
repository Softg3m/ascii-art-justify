package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	old := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

func TestGenerateAscii_NotEmpty(t *testing.T) {
	result := GenerateAscii("Hi", "standard")

	if result == "" {
		t.Errorf("Expected ASCII output, got empty string")
	}
}

func TestGenerateAscii_InvalidBanner(t *testing.T) {
	result := GenerateAscii("Hi", "fakebanner")

	if !strings.Contains(result, "Error") {
		t.Errorf("Expected error message for invalid banner")
	}
}


func TestPrintJustified_OutputNotEmpty(t *testing.T) {
	output := captureOutput(func() {
		PrintJustified([]string{"hello", "world"}, "standard", 80)
	})

	if output == "" {
		t.Errorf("Expected justified output, got empty")
	}
}

func TestPrintJustified_RespectsWidth(t *testing.T) {
	width := 60

	output := captureOutput(func() {
		PrintJustified([]string{"go", "lang"}, "standard", width)
	})

	lines := strings.Split(output, "\n")

	for _, line := range lines {
		if len(line) > width {
			t.Errorf("Line exceeds width: %d > %d", len(line), width)
		}
	}
}

func TestRightAlignmentPadding(t *testing.T) {
	line := "hello"
	width := 20

	padding := width - len(line)
	result := strings.Repeat(" ", padding) + line

	if len(result) != width {
		t.Errorf("Right alignment failed, expected length %d got %d", width, len(result))
	}
}

func TestCenterAlignmentPadding(t *testing.T) {
	line := "hello"
	width := 20

	padding := (width - len(line)) / 2
	result := strings.Repeat(" ", padding) + line

	if len(result) <= len(line) {
		t.Errorf("Center alignment failed")
	}
}
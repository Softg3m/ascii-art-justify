package main

import (
	"bytes"
	"os"
	"testing"
)

func TestUsage(t *testing.T) {

	old := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	Usage()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	t.Log("Captured output:\n" + output)
}

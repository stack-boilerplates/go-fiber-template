package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = originalStdout

	expected := "Hello, World!\n"
	if buf.String() != expected {
		t.Errorf("main() output = %v, want %v", buf.String(), expected)
	}
}

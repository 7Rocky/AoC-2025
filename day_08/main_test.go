package main

import (
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	want := "Three largest circuit sizes product (1): 175440\n" +
		"Last two junction boxes x-coordinate product (2): 3200955921\n"

	if string(out) != want {
		t.Errorf("\nWant:\n%s\nGot:\n%s", want, out)
	}
}

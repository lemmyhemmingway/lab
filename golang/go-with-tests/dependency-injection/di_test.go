package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Gundi")

	got := buffer.String()
	want := "Hello, Gundi"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

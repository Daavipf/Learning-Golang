package dinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer, "Davi")

	got := buffer.String()
	want := "Hello, Davi"

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}

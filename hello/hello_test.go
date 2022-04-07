package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Antonio")
	want := "Hello, Antonio"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

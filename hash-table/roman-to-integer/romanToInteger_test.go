package main

import "testing"

func TestRomanToInt(t *testing.T) {
	input := "XIV"
	got := romanToInt(input)
	want := 14
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

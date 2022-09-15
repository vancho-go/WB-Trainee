package rle

import (
	"testing"
)

func TestExtract(t *testing.T) {
	var tests = []struct {
		input, want string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", "некорректная строка"},
		{"", ""},
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			answer := Extract(tt.input)
			if answer != tt.want {
				t.Errorf("got %s, want %s", answer, tt.want)
			}
		})
	}
}

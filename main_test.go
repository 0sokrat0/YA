package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {

	tests := []struct {
		input    []byte
		expected int
		err      error
	}{
		{[]byte("hello"), 5, nil},
		{[]byte("Привет"), 6, nil},
		{[]byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8},
		{[]byte(""), 0, nil},
	}

	for _, tt := range tests {
		result, err := GetUTFLength(tt.input)
		if result != tt.expected || err != tt.err {
			t.Errorf("GetUTFLength(%q) = %d, %v; expected %d, %v", tt.input, result, err, tt.expected, tt.err)
		}
	}
}

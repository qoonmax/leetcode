package main

import (
	"testing"
)

func TestGenerateSlice(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"EmptySlice", 0},
		{"SingleElement", 1},
		{"MultipleElements", 5},
		{"LargeSlice", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateSlice(tt.n)

			// Check length
			if len(got) != tt.n {
				t.Errorf("generateSlice() length = %v, want %v", len(got), tt.n)
			}

			// Check uniqueness
			m := make(map[int]struct{})
			for _, v := range got {
				if _, ok := m[v]; ok {
					t.Errorf("generateSlice() contains duplicate elements")
				}
				m[v] = struct{}{}
			}
		})
	}
}

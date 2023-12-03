package main

import (
	"testing"
)

func TestIsDigit(t *testing.T) {
	tests := []struct {
		name  string
		want  bool
		input rune
		err   bool
	}{
		{"valid test integer", true, '1', false},
		{"valid test string", false, 'n', false},
		{"want error", false, '1', true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isDigit(tt.input)
			// if we receive an error, and didn't want to
			if got != tt.want && !tt.err {
				t.Errorf("wanted: %t, received %t", tt.want, got)
			}
		})
	}
}

func TestHasDigit(t *testing.T) {
	tests := []struct {
		name  string
		want  bool
		input string
		err   bool
	}{
		{"digit from string", true, "oneljfsljp", false},
		{"no digit", false, "fslifjohfwo", false},
		{"no digit and too short", false, "ij", false},
		{"integer in the string", false, "3ioj1", false},
		{"invalid check, integer in the string", true, "o3ne", true},
		{"invalid check, no digit", true, "oiwhonehois", true},
		{"invalid check, digit present", false, "oneifnwoeifnowf", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := hasDigit(tt.input)
			if got != tt.want && !tt.err {
				t.Fatalf("Error: wanted %t, got %t", tt.want, got)
			}
		})
	}
}

func TestStringToInts(t *testing.T) {
	tests := []struct {
		name  string
		want  []int
		input string
		err   bool
	}{
		{"string and int", []int{1, 2}, "one2", false},
		{"only strings", []int{1, 1, 2}, "elijonefiwoihfowonetwo", false},
		{"only ints", []int{1, 2}, "j1joihow2jfsihfos", false},
		{"invalid, int and string", []int{1, 2}, "oiwoneoneone2", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringToInts(tt.input)
			valid := true
			// length validation
			if len(got) != len(tt.want) && !tt.err {
				t.Fatalf("slices do not match, received %v, wanted %v", got, tt.want)
			} else if len(got) != len(tt.want) && tt.err {
				/*
					this is a special case where there is a mismatch in length of slices,
					but we are checking  purposely for an error
				*/
				return
			}
			// validation of each integer
			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					valid = false
				}
			}
			if !valid && !tt.err {
				t.Fatalf("received %v, wanted %v", got, tt.want)
			}
		})
	}
}

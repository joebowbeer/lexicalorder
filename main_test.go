package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestCeilLog2(t *testing.T) {
	tests := []struct {
		items int
		steps int
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 3},
		{10, 4},
		{16, 4},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.items), func(t *testing.T) {
			if got := ceilLog2(tt.items); got != tt.steps {
				t.Errorf("ceilLog2(%d) = %d; want %d", tt.items, got, tt.steps)
			}
		})
	}
}

func TestLexicalOrder(t *testing.T) {
	words := []string{"bca", "aaa", "acb"}
	expect := []string{"b", "a", "c"}
	if got := LexicalOrder(words); !reflect.DeepEqual(got, expect) {
		t.Errorf("LexicalOrder returned %v; want %v", got, expect)
	}
}

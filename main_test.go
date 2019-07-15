package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestLexicalOrder(t *testing.T) {
	tests := []struct {
		name   string
		words  []string
		expect []string
	}{
		{"(empty)",
			nil,
			nil},
		{"a",
			[]string{"a", "aa"},
			[]string{"a"}},
		{"ab",
			[]string{"a", "b"},
			[]string{"a", "b"}},
		// Test cases from:
		// https://www.geeksforgeeks.org/given-sorted-dictionary-find-precedence-characters/
		{"bac",
			[]string{"bca", "aaa", "acb"},
			[]string{"b", "a", "c"}},
		{"bdac",
			[]string{"baa", "abcd", "abca", "cab", "cad"},
			[]string{"b", "d", "a", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LexicalOrder(tt.words); !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("LexicalOrder returned %v; want %v", got, tt.expect)
			}
		})
	}
}

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

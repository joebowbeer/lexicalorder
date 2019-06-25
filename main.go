package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
)

var numChunks = runtime.GOMAXPROCS(0) // TODO: best practice?

// Reads an ordered list of words from stdin and prints the determined character sort order.
func main() {
	fmt.Println(LexicalOrder(readInput(*bufio.NewScanner(os.Stdin))))
}

/*LexicalOrder receives a list of words that are sorted according to an unknown
character precedence and returns their characters in the determined order.*/
func LexicalOrder(words []string) []string {
	slices, runes := indexRunes(words)
	dist := adjacency(slices, len(runes))
	for n := ceilLog2(len(dist)); n > 0; n-- {
		dist = maxplus(dist)
	}
	return restoreRunes(sortedIndices(dist), runes)
}

// Returns list of words read from scanner.
func readInput(scanner bufio.Scanner) (words []string) {
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

// Given a list of words, converts each rune to an index into a list of runes.
// Returns the indices and the indexed list of runes.
func indexRunes(words []string) (slices [][]int, runes []rune) {
	m := map[rune]int{}
	for _, word := range words {
		indices := []int{}
		for _, r := range word {
			if _, ok := m[r]; !ok {
				m[r] = len(runes)
				runes = append(runes, r)
			}
			indices = append(indices, m[r])
		}
		slices = append(slices, indices)
	}
	return
}

// Given a list of indices and an indexed list of runes, returns list of single-rune strings.
func restoreRunes(indices []int, runes []rune) (chars []string) {
	for _, index := range indices {
		chars = append(chars, string(runes[index]))
	}
	return
}

// Returns adjacency matrix given sorted list of slices containing ints in range [0..dim).
func adjacency(in [][]int, dim int) [][]int {
	// initialize output
	out := make([][]int, dim)
	for i := range out {
		out[i] = make([]int, dim)
	}
	for i := range in {
		if i == 0 {
			continue
		}
		// pred precedes succ in lexical order
		pred := in[i-1]
		succ := in[i]
		// find first position where their ints differ
		for k := 0; k < len(pred) && k < len(succ); k++ {
			if pred[k] != succ[k] {
				out[pred[k]][succ[k]] = 1
				break
			}
		}
	}
	return out
}

// Computes ceil(log2(dim))
func ceilLog2(dim int) int {
	n := 0
	// pow is 2^n
	for pow := 1; pow < dim; pow <<= 1 {
		n++
	}
	return n
}

// Returns max-plus product of the given distance matrix.
func maxplus(in [][]int) [][]int {
	dim := len(in)
	out := make([][]int, dim)
	var wg sync.WaitGroup
	chunkSize := (dim + numChunks - 1) / numChunks
	for i := 0; i < dim; i += chunkSize {
		end := i + chunkSize
		if end > dim {
			end = dim
		}
		wg.Add(1)
		go func(m, n int) {
			defer wg.Done()
			for i := m; i < n; i++ {
				out[i] = maxplusRow(in, i)
			}
		}(i, end)
	}
	wg.Wait()
	return out
}

// Returns one row of the max-plus product of the given distance matrix.
func maxplusRow(in [][]int, i int) []int {
	dim := len(in)
	row := make([]int, dim)
	for j := 0; j < dim; j++ {
		if i == j {
			continue
		}
		max := in[i][j]
		for k := 0; k < dim; k++ {
			if in[i][k] == 0 || in[k][j] == 0 {
				continue
			}
			if sum := in[i][k] + in[k][j]; sum > max {
				max = sum
			}
		}
		row[j] = max
	}
	return row
}

// Returns row indices ordered by the longest distance found in each row.
func sortedIndices(dist [][]int) []int {
	dim := len(dist)
	out := make([]int, dim)
	for i := range out {
		out[i] = -1
	}
	for index := range dist {
		// find maximum distance in row
		max := 0
		for _, d := range dist[index] {
			if d > max {
				max = d
			}
		}
		if max >= dim {
			log.Fatalf("Fail! Cycle detected at index %d", index)
		}
		rank := dim - max - 1
		if out[rank] != -1 {
			log.Fatalf(
				"Fail! Rank %d of index %d already assigned to index %d",
				rank, index, out[rank])
		}
		out[rank] = index
	}
	return out
}

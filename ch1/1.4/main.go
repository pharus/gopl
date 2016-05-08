// pharus solved this @ Mon 25 Apr 16:49:23 WEST 2016

// Exercise 1.4 from page 13

// moded dup2 to print the name of files the duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		counts["os.Stdin"] = make(map[string]int)
		countLines(os.Stdin, counts["os.Stdin"])
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
				continue
			}

			counts[arg] = make(map[string]int)
			countLines(f, counts[arg])
			f.Close()
		}
	}

	for file, count := range counts {
		for line, n := range count {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", file, n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

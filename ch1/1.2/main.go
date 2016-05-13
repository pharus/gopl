// pharus solved this @ Mon 25 Apr 03:47:21 WEST 2016

// Exercise 1.2 from page 8

// echo program to print the name of the program too
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}

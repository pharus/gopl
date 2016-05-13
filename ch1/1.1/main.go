// pharus solved this @ Mon 25 Apr 03:44:16 WEST 2016

// Exercise 1.1 from page 8

// echo program to print the name of the program too
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}

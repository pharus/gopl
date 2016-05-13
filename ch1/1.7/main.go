// pharus solved this @ Wed 27 Apr 00:05:48 WEST 2016

// Exercise 1.7 from page 17

// Fetch using ioutil.Copy
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: reading %s: %v\n", os.Args[0], url, err)
			os.Exit(1)
		}
	}
}

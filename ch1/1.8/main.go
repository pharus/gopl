// pharus solved this @ Wed 27 Apr 00:15:29 WEST 2016

// Exercise 1.8 from page 17

// Adding prefix if absent
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
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

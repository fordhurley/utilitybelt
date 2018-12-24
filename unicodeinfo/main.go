package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/unicode/runenames"
)

const usage = `unicodeinfo

Prints out unicode information for each character on stdin, or from the
provided arguments.

Usage: unicodeinfo [-h | "String to decode!"]
	-h  print this help
`

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}

	var input io.Reader = os.Stdin
	if len(os.Args) > 1 {
		input = strings.NewReader(strings.Join(os.Args[1:], " "))
	}
	b := bufio.NewReader(input)

	var index int
	var r rune
	var size int
	var err error

	for {
		r, size, err = b.ReadRune()
		if err != nil {
			break
		}
		fmt.Printf("%d: %#U (%s)\n", index, r, runenames.Name(r))
		index += size
	}

	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "Invalid input:", err)
		os.Exit(2)
	}
}

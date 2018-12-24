package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	var err error

	for {
		r, _, err = b.ReadRune()
		if err != nil {
			break
		}
		fmt.Printf("%d: %#U\n", index, r)
		index++
	}

	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "Invalid input:", err)
		os.Exit(2)
	}
}

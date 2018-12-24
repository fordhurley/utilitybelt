package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	var input io.Reader = os.Stdin
	if len(os.Args) > 1 {
		var err error
		input, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	counter := Counter(0)
	w := gzip.NewWriter(&counter)
	written, err := io.Copy(w, input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = w.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	originalSize := int(written)
	compressedSize := int(counter)
	factor := float64(compressedSize) / float64(originalSize)

	fmt.Printf(
		"%s -> %s (%.03f)\n",
		FormatSize(originalSize),
		FormatSize(compressedSize),
		factor,
	)
}

// FormatSize formats a file size, like "1.234 kB".
//
// How many times has this been written before? Doesn't feel like it's worth a
// dependency right now.
func FormatSize(n int) string {
	prefixes := []string{"", "k", "M", "G"}

	index := 0
	nF := float64(n)
	for nF > 800 && index < len(prefixes)-1 {
		nF /= 1000
		index++
	}

	// intentionally losing precision beyond 3 digits after decimal
	nF = math.Trunc(nF*1000) / 1000

	return fmt.Sprintf("%v %sB", nF, prefixes[index])
}

// Counter implements io.Writer, to count the number of bytes written to it,
// discarding the actual bytes.
type Counter int

func (c *Counter) Write(p []byte) (int, error) {
	n := len(p)
	*c += Counter(n)
	return n, nil
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n    %s Здравствуй, мир!\n", os.Args[0])
		os.Exit(1)
	}

	input := strings.Join(os.Args[1:], " ")

	for index, runeValue := range input {
		fmt.Printf("%d: %#U\n", index, runeValue)
	}
}

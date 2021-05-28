package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"math"
	"math/big"
	"os"
)

func main() {
	cmd := "uint32"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	r := randy{rand.Reader}

	switch cmd {
	case "uint32":
		fmt.Println(r.Uint32())
	default:
		fmt.Fprintf(os.Stderr, "ERROR: unknown command %q\n", cmd)
	}
}

type randy struct {
	io.Reader
}

func (r randy) Uint32() uint32 {
	n, err := rand.Int(r, big.NewInt(math.MaxUint32))
	if err != nil {
		panic(err)
	}
	return uint32(n.Uint64())
}

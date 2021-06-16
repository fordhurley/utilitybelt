package main

import (
	"crypto/rand"
	"encoding/base64"
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
	case "uuid":
		fmt.Println(r.UUID())
	case "token":
		fmt.Println(r.Token())
	default:
		fmt.Fprintf(os.Stderr, "ERROR: unknown command %q\n", cmd)
	}
}

type randy struct {
	io.Reader
}

func (r randy) Uint32() *big.Int {
	n, err := rand.Int(r, big.NewInt(math.MaxUint32))
	if err != nil {
		panic(err)
	}
	return n
}

func (r randy) UUID() string {
	bs := make([]byte, 16)
	_, err := io.ReadFull(r, bs)
	if err != nil {
		panic(err)
	}
	bs[6] = (bs[6] & 0x0f) | 0x40 // version=4
	bs[8] = (bs[8] & 0x3f) | 0x80 // variant=10

	return fmt.Sprintf("%x-%x-%x-%x-%x", bs[:4], bs[4:6], bs[6:8], bs[8:10], bs[10:])
}

func (r randy) Token() string {
	bs := make([]byte, 48) // 48 bytes, base64 encoded => 64 character string
	_, err := io.ReadFull(r, bs)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(bs)
}

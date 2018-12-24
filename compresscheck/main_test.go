package main

import "fmt"

func ExampleFormatSize() {
	fmt.Println(FormatSize(1))
	fmt.Println(FormatSize(10))
	fmt.Println(FormatSize(900))
	fmt.Println(FormatSize(999))
	fmt.Println(FormatSize(1000))
	fmt.Println(FormatSize(1042))
	fmt.Println(FormatSize(999999))
	fmt.Println(FormatSize(1000000))
	fmt.Println(FormatSize(1000042))
	fmt.Println(FormatSize(1042000))

	// Output:
	// 1 B
	// 10 B
	// 0.9 kB
	// 0.999 kB
	// 1 kB
	// 1.042 kB
	// 0.999 MB
	// 1 MB
	// 1 MB
	// 1.042 MB
}

package main

import (
	"fmt"

	prettybytes "github.com/fernandoporazzi/pretty-bytes"
)

func main() {
	fmt.Println(prettybytes.Convert(10.1))
	// => 10 B

	fmt.Println(prettybytes.Convert(10, prettybytes.Options{
		Signed: true,
	}))
	// => +10 B

	fmt.Println(prettybytes.Convert(98765, prettybytes.Options{
		Bits: true,
	}))
	// => 99 kbit

	fmt.Println(prettybytes.Convert(98765, prettybytes.Options{
		Signed: true,
		Bits:   true,
	}))
	// => +99 kbit
}

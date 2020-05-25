# pretty-bytes
Convert bytes to a human readable string: 1337 → 1.34 kB

## Install

```
$ go get -u github.com/fernandoporazzi/pretty-bytes
```

## API

```go
func Convert(n float64, options ...Options) string
```

```go
type Options struct {
	// Format the number as bits instead of bytes.
	// This can be useful when, for example, referring to bit rate.
	// Defaults to false
	Bits bool

	// Include plus sign for positive numbers.
	// Defaults to false
	Signed bool
}
```


## Usage

```go
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
```

## Related
- [pretty-milliseconds](https://github.com/fernandoporazzi/pretty-milliseconds) - Convert milliseconds to a human readable string: 1337000000 → 15d 11h 23m 20s

## License
[MIT license](https://github.com/fernandoporazzi/pretty-bytes/blob/master/LICENSE)

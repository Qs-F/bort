# pkg `bort`

This package classify given data into binary and text (UTF-8 only)

## Installation

```
go get github.com/Qs-F/bort
```

## Usage 

Example to classify data from stdin into binary data and text data

```go
package main

import (
	"fmt"
	"os"

	"github.com/Qs-F/bort"
)

func main() {
	b, err := bort.IsBin(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	if b {
		fmt.Println("binary")
	} else {
		fmt.Println("text")
	}
}
```

## License

MIT License

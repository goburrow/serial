# serial [![GoDoc](https://godoc.org/github.com/goburrow/serial?status.svg)](https://godoc.org/github.com/goburrow/serial)
For Windows users, a gcc compiler is required. See [cgo](https://github.com/golang/go/wiki/cgo#windows) for more information.
## Example
```go
package main

import (
	"log"

	"github.com/goburrow/serial"
)

func main() {
	port, err := serial.Open(&serial.Config{Address: "/dev/ttyUSB0"})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	_, err = port.Write([]byte("serial"))
	if err != nil {
		log.Fatal(err)
	}
}
```
## Testing

### Linux
- `socat -d -d pty,raw,echo=0 pty,raw,echo=0`

### Windows
- [Null-modem emulator](http://com0com.sourceforge.net/)
- [Terminal](https://sites.google.com/site/terminalbpp/)

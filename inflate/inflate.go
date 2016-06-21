package main

import (
	"compress/flate"
	"io"
	"os"
)

func main() {
	r := flate.NewReader(os.Stdin)
	io.Copy(os.Stdout, r)
	r.Close()
}

package main

import (
	"compress/flate"
	"io"
	"os"
)

func main() {
	w, err := flate.NewWriter(os.Stdout, 1)
	if err != nil {
		panic(err)
	}
	io.Copy(w, os.Stdin)
	w.Close()
}

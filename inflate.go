package main

import (
	"compress/flate"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var d = flag.Bool("d", false, "Deflates input instead")
var l = flag.Int("l", -1, "Deflate compression level")
var alg = flag.String("alg", "flate", "The compression algorithm to use")

func main() {
	flag.Parse()
	if *d {
		w, err := NewWriter()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(w, os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		err = w.Close()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		r, err := NewReader()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, r)
		if err != nil {
			log.Fatal(err)
		}
		err = r.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func NewWriter() (io.WriteCloser, error) {
	switch *alg {
	case "flate":
		return flate.NewWriter(os.Stdout, *l)
	default:
		return nil, fmt.Errorf("Unknown Algorithm")
	}
}

func NewReader() (io.ReadCloser, error) {
	switch *alg {
	case "flate":
		return flate.NewReader(os.Stdin), nil
	default:
		return nil, fmt.Errorf("Unknown Algorithm")
	}
}

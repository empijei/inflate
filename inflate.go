package main

import (
	"compress/flate"
	"compress/zlib"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var d = flag.Bool("d", false, "Deflates input instead")
var l = flag.Int("l", -1, "Deflate compression level")
var alg = flag.String("alg", "flate", "The compression algorithm to use.\n\tSupported algorithms:\n\t\t'flate','f'\n\t\t'zlib','z'\n")

func main() {
	flag.Parse()
	if *d {
		w, err := NewWriter()
		errFatal(err)
		_, err = io.Copy(w, os.Stdin)
		errFatal(err)
		err = w.Close()
		errFatal(err)
	} else {
		r, err := NewReader()
		errFatal(err)
		_, err = io.Copy(os.Stdout, r)
		errFatal(err)
		err = r.Close()
		errFatal(err)
	}
}

/*TODO
https://golang.org/pkg/compress/gzip/
https://golang.org/pkg/compress/lzw/
https://golang.org/pkg/compress/bzip2/
*/

func NewWriter() (io.WriteCloser, error) {
	switch *alg {
	case "flate", "f":
		return flate.NewWriter(os.Stdout, *l)
	case "zlib", "z":
		return zlib.NewWriter(os.Stdout), nil
	default:
		return nil, fmt.Errorf("Unknown Algorithm")
	}
}

func NewReader() (io.ReadCloser, error) {
	switch *alg {
	case "flate", "f":
		return flate.NewReader(os.Stdin), nil
	case "zlib", "z":
		return zlib.NewReader(os.Stdin)
	default:
		return nil, fmt.Errorf("Unknown Algorithm")
	}
}

func errFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

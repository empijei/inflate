package main

import (
	"compress/bzip2"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var d = flag.Bool("d", false, "Compresses input instead")
var l = flag.Int("level", -1, "Deflate compression level")
var alg = flag.String("alg", "flate", `The compression algorithm to use
	Supported algorithms: 
		'flate','f' 
		'zlib' ,'z' 
		'gzip' ,'g'`)

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
https://golang.org/pkg/compress/lzw/
https://golang.org/pkg/compress/bzip2/
*/

func NewWriter() (io.WriteCloser, error) {
	switch *alg {
	case "flate", "f":
		return flate.NewWriter(os.Stdout, *l)
	case "zlib", "z":
		return zlib.NewWriter(os.Stdout), nil
	case "gzip", "g":
		return gzip.NewWriter(os.Stdout), nil
	case "bzip2", "b":
		return nil, fmt.Errorf("Bzip does not support compressing yet.")
	default:
		return nil, fmt.Errorf("Unknown Algorithm")
	}
}

type bzip2Wrapper struct {
	io.Reader
}

func (b bzip2Wrapper) Close() error {
	return nil
}

func NewReader() (io.ReadCloser, error) {
	switch *alg {
	case "flate", "f":
		return flate.NewReader(os.Stdin), nil
	case "zlib", "z":
		return zlib.NewReader(os.Stdin)
	case "gzip", "g":
		return gzip.NewReader(os.Stdin)
	case "bzip2", "b":
		b := bzip2.NewReader(os.Stdin)
		toret := bzip2Wrapper{b}
		return toret, nil
	default:
		return nil, fmt.Errorf("Unknown Algorithm")
	}
}

func errFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

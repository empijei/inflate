package main

import (
	"compress/flate"
	"flag"
	"io"
	"log"
	"os"
)

var d = flag.Bool("d", false, "Deflates input instead")

func main() {
	flag.Parse()
	if *d {
		w, err := flate.NewWriter(os.Stdout, 1)
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
		r := flate.NewReader(os.Stdin)
		_, err := io.Copy(os.Stdout, r)
		if err != nil {
			log.Fatal(err)
		}
		err = r.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

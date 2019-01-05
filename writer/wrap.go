package main

import (
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please give file name.")
		return
	}
	fn := os.Args[1]

	f, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	// Declare writer variable
	var writer io.Writer

	// Assign os.Stdout as io.Writer
	writer = os.Stdout

	// Wrap writer by base64.Encoder
	encoder := base64.NewEncoder(base64.StdEncoding, writer)
	defer encoder.Close()
	writer = encoder

	// Wrap writer by gzip
	writer = gzip.NewWriter(writer)

	io.Copy(writer, f)
}

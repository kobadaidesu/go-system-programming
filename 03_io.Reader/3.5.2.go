package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("example.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	gz, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer gz.Close()
	io.Copy(os.Stdout, gz)
}
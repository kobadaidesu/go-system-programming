//go:build ignore

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("strings.Builder example\n"))
	fmt.Println(buffer.String())
}

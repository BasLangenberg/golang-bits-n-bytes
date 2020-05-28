package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fp, err := os.Open("./resources.go")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	b := make([]byte, 1)
	for {
		n, err := fp.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		fmt.Printf("%s", string(b[:n]))
	}
}

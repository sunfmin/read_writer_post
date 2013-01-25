package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	var r io.ReadSeeker
	var err error
	r, err = os.Open("/tmp/hello.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(findGolang(r))
}

func findGolang(r io.Reader) (position int) {
	var err error

	var G = []byte("G")
	var olang = []byte("olang")

	currentByte := make([]byte, 1)
	readByte := 0

	var l int
	for {
		l, err = r.Read(currentByte)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		readByte += l

		if bytes.Compare(currentByte, G) == 0 {
			next5 := make([]byte, 5)
			l, err = r.Read(next5)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			if bytes.Compare(next5, olang) == 0 {
				return readByte - 1
			}
		}
	}
	return -1
}

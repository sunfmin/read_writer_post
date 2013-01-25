package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	var r io.Reader
	var err error
	r, err = os.Open("/tmp/hello.txt")
	if err != nil {
		panic(err)
	}
	var body []byte
	body, err = ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

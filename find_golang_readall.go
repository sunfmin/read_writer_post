package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	var r io.Reader
	var err error
	r, err = os.Open("/tmp/hello.txt")

	var body []byte
	body, err = ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	found := strings.Index(string(body), "Golang")
	fmt.Println(found)
}

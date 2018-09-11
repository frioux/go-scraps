package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dir, err := ioutil.TempDir("/home/frew/tmp/x", "secrets")
	if err != nil {
		panic(err)
	}
	defer os.Remove(dir)
	fmt.Println(dir)
}

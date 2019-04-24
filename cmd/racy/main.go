package main

import "fmt"

func main() {
	x := map[string]string{"frew": "1"}

	go func() {
		for i := 0; ; i++ {
			x = map[string]string{"frew": fmt.Sprintf("%d", i)}
		}
	}()

	for {
		fmt.Println(x["frew"])
	}
}

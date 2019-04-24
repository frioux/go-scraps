package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {

	go func() {
		for {
			time.Sleep(time.Millisecond)
			if rand.Intn(2) == 0 {
				os.Exit(1)
			}
		}
	}()

	for {
		time.Sleep(time.Millisecond)
	}
}

package main

import (
	"flag"
	"fmt"
	"time"
)

var duration time.Duration

func init() {
	flag.DurationVar(&duration, "duration", time.Millisecond, "how long to sleep")
}

func main() {
	flag.Parse()

	for {
		start := time.Now()
		timer := time.NewTimer(duration)

		<-timer.C
		end := time.Now()
		fmt.Println(end.Sub(start) - duration)
	}
}

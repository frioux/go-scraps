package main

import "fmt"

func main() {
	x := []int{1}
	m := x[0]
	for _, v := range x[1:] {
		if v < m {
			m = v
		}
	}
	fmt.Println(m)
}

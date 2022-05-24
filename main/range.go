package main

import (
	"fmt"
)

func main() {
	var c = []int{1, 2, 3, 4, 5}
	for idx, num := range c {
		fmt.Println(idx, "-", num)
	}
	m := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range m {
		fmt.Println(k, "--", v)
	}

}

package main

import (
	"fmt"
)

func main() {
	var a [5]int

	fmt.Println("emp:", a)

	a[4] = 500
	fmt.Println("set:", a)
	fmt.Println("get:", a)

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 3, 4}
	fmt.Println("dcl=:", b)

	var c = new([10]int)
	s := append(c[:], 2)
	fmt.Println(s)

	var twoD [2][4]int
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)
}

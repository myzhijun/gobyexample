package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"a", "b", "c"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{1, 2, 1, 3, 5, 2}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints) //
	fmt.Println("Sorted:", s)
}

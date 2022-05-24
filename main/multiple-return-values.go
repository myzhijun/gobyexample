package main

import "fmt"

func addmul(a ...int) (addAns int, mulAns int) {
	if len(a) == 0 {
		return 0, 0
	}
	addAns = 0
	mulAns = 1
	for _, n := range a {
		addAns += n
		mulAns *= n
	}
	return
}
func main() {
	addAns, mulAns := addmul()
	fmt.Println("ans:", addAns, ",", mulAns)
}

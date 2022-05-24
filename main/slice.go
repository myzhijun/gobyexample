package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "A"
	s[1] = "B"
	s[2] = "C"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sli:", l)

	l = s[:5]
	fmt.Println("sli2:", l)
}

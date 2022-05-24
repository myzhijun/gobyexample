package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"bob", 12})
	fmt.Println(person{name: "alice", age: 13})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "ann", age: 34})
	s := person{"Sean", 23}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)
	ss := s
	ss.age = 10
	fmt.Println(ss.age, "/", s.age)
}

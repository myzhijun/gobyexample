package main

import "fmt"

type Rect struct {
	width, height float64
}

func (r *Rect) changeptr() {
	r.width = 100
}
func (r Rect) changevalue() {
	r.width = 100
}
func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perim() float64 {
	return 2 * (r.width + r.height)
}
func main() {
	r := Rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())
	fmt.Println("width:", r.width)
	r.changevalue()
	fmt.Println("width:", r.width)

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
	fmt.Println("width:", r.width)
	rp.changevalue()
	fmt.Println("width:", rp.width)

}

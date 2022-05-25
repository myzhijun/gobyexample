package main

import (
	"fmt"
	"os"
)

func main() {
	f := creatFile("/tmp/defer.txt")
	defer closeFile(f)
	writerFile(f)
}
func creatFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func readfile(f *os.File) {
	fmt.Println("reading")
}
func writerFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

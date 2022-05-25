package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOF", "1")
	fmt.Println("FOF", os.Getenv("FOF"))
	fmt.Println("BOB", os.Getenv("BOB"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], ":", pair[1])
	}

}

package main

import (
	"fmt"
	S "strings"
)

var p = fmt.Println

func main() {
	p("Contains:", S.Contains("test", "es"))
	p("Count:", S.Count("test", "t"))
	p("HasPrefix:", S.HasPrefix("test", "te"))
	p("HasSuffix:", S.HasSuffix("test", "e"))
	p("Index:", S.Index("test", "e"))
	p("Join:", S.Join([]string{"a", "b"}, "-"))
	p("Repeat:", S.Repeat("a", 5))
	p("Replace:", S.Replace("foo", "o", "0", -1))
	p("Split:", S.Split("a-b-c-d-e", "-"))
	p("ToLower:", S.ToLower("TEST"))
	p("ToUpper:", S.ToUpper("test"))
	p()
	p("Len:", len("hello"))
	p("Char:", "hello"[1])
}

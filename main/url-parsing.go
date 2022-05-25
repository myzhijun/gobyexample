package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	s := "https://fanyi.baidu.com/?aldtype=16047#en/zh/scheme"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("user:", u.User)
	fmt.Println("username:", u.User.Username())
	password, _ := u.User.Password()
	if err != nil {
		panic(err)
	}
	fmt.Println("password:", password)
	fmt.Println("host:", u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	//fmt.Println(h[1])
	fmt.Println("path:", u.Path)
	fmt.Println("fragment:", u.Fragment)
	fmt.Println("rawQuery:", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	//fmt.Println(m["k"][0])
}

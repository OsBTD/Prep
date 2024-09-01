package main

import (
	"fmt"
)

func Atoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	var s string
	var sign int = 1
	for i := 0; i < len(str); i++ {
		if string(str[0]) == "-" && string(str[1]) != "-" {
			sign = -1
		}
		if string(str[i]) >= "0" && string(str[i]) <= "9" {
			s += string(str[i])
		}
	}
	var c int
	var n int
	n = int(s[0]-48) * 10
	for i := 0; i < len(s)-1; i++ {

		c = n + int(s[i+1]-48)
		n = c * 10

	}

	return (n / 10) * sign

}
func main() {
	a := "12456"
	b := "12345"

	c := "-12456"
	d := "--12345"
	e := "-5"
	f := ""

	fmt.Println(Atoi(a))
	fmt.Println(Atoi(b))
	fmt.Println(Atoi(c))
	fmt.Println(Atoi(d))
	fmt.Println(Atoi(e))
	fmt.Println(Atoi(f))

}

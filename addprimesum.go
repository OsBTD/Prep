package main

import (
	"fmt"
)

func IsPrime(x int) bool {
	var isprime bool = false
	for i := x - 1; i > 1; i-- {
		if x%i == 0 {
			isprime = false
		} else {
			isprime = true
		}

	}
	return isprime

}

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	var c int
	var n int
	if rune(s[0]) <= '0' && rune(s[0]) >= '9' {
		n = int(s[0]-48) * 10
	}
	for i := 0; i < len(s)-1; i++ {
		if rune(s[i]) <= '0' && rune(s[i]) >= '9' {
			continue
		} else {

			c = n + int(s[i+1]-48)
			n = c * 10

		}

	}

	return n / 10

}
func main() {
	a := "12456"
	b := "12345"

	c := "-12456"
	d := "-12345"
	e := "-5"
	f := ""

	fmt.Println(Atoi(a))
	fmt.Println(Atoi(b))
	fmt.Println(Atoi(c))
	fmt.Println(Atoi(d))
	fmt.Println(Atoi(e))
	fmt.Println(Atoi(f))

}

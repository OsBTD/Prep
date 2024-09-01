package main

import (
	"fmt"
	"os"
)

func IsPrime(x int) bool {
	var isprime bool = false
	if x == 2 {
		isprime = true
	}
	for i := x - 1; i > 1; i-- {
		if x%i == 0 {
			isprime = false
			break
		} else {
			isprime = true
		}

	}
	return isprime

}

func Atoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	var s string
	var sign int = 1
	for i := 0; i < len(str); i++ {
		if string(str[0]) == "-" {
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
	var result int
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("0")
		return
	}
	if len(args) == 1 && Atoi(args[0]) < 0 {
		fmt.Println("0")
		return
	}
	for i := Atoi(args[0]); i > 1; i-- {
		if IsPrime(i) {
			fmt.Println(i, IsPrime(i))
			result += i
		}

	}

	fmt.Println(result)
}

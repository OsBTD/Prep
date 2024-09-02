package main

import "fmt"

func main() {
	a := 2555
	var res []string
	hexa := [16]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	for a > 0 {
		res = append(res, hexa[a%16])
		a /= 16
	}
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i])
	}
}

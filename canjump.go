package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))

	input4 := []uint{3, 2, 1, 1, 4}
	fmt.Println("basically input 2 but can indeed jump :", CanJump(input4))

	input5 := []uint{3, 2, 1, 5, 4}
	fmt.Println("out of range test :", CanJump(input5))

}

func CanJump(steps []uint) bool {
	var jump bool
	var hop int
	for i := 0; i < len(steps); i++ {
		if i < len(steps)-1 && i + int(steps[i]) < len(steps) {
			hop = i + int(steps[i])
		}
		if i == len(steps)-1 && (steps[i] == steps[hop]) {
			jump = true
		} else {
			jump = false
		}

	}
	return jump
}

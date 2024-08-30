package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func WeAreUnique(str1, str2 string) int {
	var unique []string
	if str1 == "" || str2 == "" {
		return -1
	}
	var un bool = false
	var str []string
	for i := 0; i < len(str1); i++ {
		str = append(str, string(str1[i]))
	}
	for i := 0; i < len(str2); i++ {
		str = append(str, string(str2[i]))
	}

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str); j++ {
			if i != j && str[i] == str[j] {
				un = false
				break
			} else {
				un = true
			}
		}

		if un {
			unique = append(unique, str[i])
		} else {
			continue
		}

	}
	return len(unique)
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	var count int = 1
	var result string
	for i := 0; i < len(s)-1; i++ {
		if string(s[i]) >= "A" && string(s[i]) <= "Z" {
			if s[i+1] == s[i] || s[i+1] == s[i]+32 {
				count++
			} else if s[i+1] != s[i] || s[i+1] != s[i]+32 {
				result += strconv.Itoa(count) + string(s[i])
				count = 1
			}
		} else if string(s[i]) >= "a" && string(s[i]) <= "z" {
			if s[i+1] == s[i] || s[i+1] == s[i]-32 {
				count++
			} else if s[i+1] != s[i] || s[i+1] != s[i]-32 {
				result += strconv.Itoa(count) + string(s[i])
				count = 1
			}
		}
	}
	return result
}

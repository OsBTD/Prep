package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println()
	}

	fmt.Println(Union(args[0], args[1]))
}

func Union(a, b string) string {
	var s []string
	var result []string
	var res string
	// var rep bool = false

	for i := 0; i < len(a); i++ {
		s = append(s, string(a[i]))
	}
	for i := 0; i < len(b); i++ {
		s = append(s, string(b[i]))
	}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				s[j] = ""
				continue
			}
		}
		result = append(result, s[i])
	}
	for i := 0; i < len(result); i++ {
		if result[i] != "" {
			res += result[i]
		}
	}
	return res
}

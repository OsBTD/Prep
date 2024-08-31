package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal("couldn't read file")
	}

	Lines := strings.Split(string(content), "\n")
	Replace := make(map[int][]string)
	Char := 1
	for i := 0; i < len(Lines); i += 9 {
		if i+9 <= len(Lines)-1 {
			Replace[Char] = Lines[i+1 : i+9]
		}
		if Char <= 9 {
			Char++
		}

	}
	for i := 9; i >= 1; i-- {
		clearScreen()
		for _, line := range Replace[i] {
			fmt.Println(line)
		}
		time.Sleep(1 * time.Second)
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

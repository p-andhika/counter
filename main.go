package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	wordCount := 0

	for _, x := range data {
		if x == ' ' {
			wordCount += 1
		}
	}

	wordCount++

	fmt.Println(wordCount)

	fmt.Println("data:", string(data))
}

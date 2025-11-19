package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	file, err := os.Open("./words.txt")
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	PrintFileContents(file)
}

func PrintFileContents(file *os.File) {
	const bufferSize = 8192
	buffer := make([]byte, bufferSize)

	totalSize := 0
	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}

		totalSize += size
		// fmt.Print(string(buffer[:size]))
	}

	fmt.Println("total bytes read:", totalSize)
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)

	return len(words)
}

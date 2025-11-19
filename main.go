package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	log.SetFlags(0)

	file, err := os.Open("./words.txt")
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	wordCount := CountWordsInFile(file)

	fmt.Println(wordCount)
}

func CountWordsInFile(file *os.File) int {
	wordCount := 0
	isInsideWord := false

	const bufferSize = 1
	buffer := make([]byte, bufferSize)
	leftover := []byte{}

	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}

		subBuffer := append(leftover, buffer[:size]...)

		for len(subBuffer) > 0 {
			r, rSize := utf8.DecodeRune(subBuffer)

			if r == utf8.RuneError {
				break
			}

			subBuffer = subBuffer[rSize:]

			if !unicode.IsSpace(r) && !isInsideWord {
				wordCount++
			}

			isInsideWord = !unicode.IsSpace(r)
		}

		leftover = nil
		leftover = append(leftover, subBuffer...)

	}

	return wordCount
}

func CountWords(data []byte) int {
	words := bytes.Fields(data)

	return len(words)
}

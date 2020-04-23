package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	pattern := os.Args[1]
	filename := os.Args[2]

	searchInFile(pattern, filename)
}

func searchInFile(pattern string, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "")

		for i := 0; i <= len(tokens)-len(pattern); i++ {
			if line[i:i+len(pattern)] == pattern {
				fmt.Println(line)
			}
		}

	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

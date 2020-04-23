package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maxCapacity = 1000

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	pattern := os.Args[1]
	filename := os.Args[2]

	searchInFile(pattern, filename)
}

func searchInFile(pattern string, filename string) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {

		line := scanner.Text()
		tokens := strings.Split(line, "")

		for i := 0; i <= len(tokens)-len(pattern); i++ {
			if line[i:i+len(pattern)] == pattern {
				fmt.Println(line, "\n")
				break
			}
		}

	}

	err = scanner.Err()
	check(err)
}

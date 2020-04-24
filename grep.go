package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const maxCapacity = 1024 * 2

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	pattern := string(os.Args[1])
	if os.Args[2] == "-i" {
		filename := os.Args[3]
		pattern = strings.ToLower(pattern)
		searchInFile(pattern, filename, true)
	} else {
		filename := os.Args[2]
		searchInFile(pattern, filename, false)
	}
}

func searchInFile(pattern string, filename string, ignore bool) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {

		line := scanner.Text()
		tokens := strings.Split(line, "")
		line1 := line

		if ignore {
			line1 = strings.ToLower(line)
		}

		for i := 0; i <= len(tokens)-len(pattern); i++ {
			if line1[i:i+len(pattern)] == pattern {
				fmt.Println(line, "\n")
				break
			}
		}

	}

	err = scanner.Err()
	check(err)
}

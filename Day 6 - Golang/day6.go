package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open file
	f, err := os.Open("day6-input")

	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println("Part 1 Answer:")
		fmt.Println(getNthUnique(string(scanner.Text()), 4))
		fmt.Println("Part 2 Answer:")
		fmt.Println(getNthUnique(string(scanner.Text()), 14))
	}

}

func unique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}

func getNthUnique(s string, n int) int {
	for i := range s {
		if i <= n-1 {
			continue
		}

		if unique(s[i-n : i]) {
			fmt.Println(s[i-n : i])
			return i
		}
	}
	return 0
}

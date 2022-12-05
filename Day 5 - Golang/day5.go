package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var m = map[int]int{
	2:  0,
	6:  1,
	10: 2,
	14: 3,
	18: 4,
	22: 5,
	26: 6,
	30: 7,
	34: 8,
}

var columns = [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}}

func main() {
	// open file
	f, err := os.Open("day5-input")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lc := 0
	for scanner.Scan() {
		lc++
		fmt.Println(scanner.Text())
		if lc > 10 {
			slice := strings.Split(scanner.Text(), " ")
			fmt.Println(slice[1], slice[3], slice[5])

			amt, _ := strconv.Atoi(slice[1])
			from, _ := strconv.Atoi(slice[3])
			to, _ := strconv.Atoi(slice[5])
			executeMove(amt, from-1, to-1)

		} else if lc == 10 || lc == 9 {
			fmt.Println("skipping")

		} else if lc < 10 {
			for i := range scanner.Text() {
				if isLetter(scanner.Text()[i]) {
					col := m[i+1]
					columns[col] = append(columns[col], string(scanner.Text()[i]))
				}
			}
		}
	}

	fmt.Println(columns)
	for i := range columns {
		print(columns[i][0])
	}
}

func executeMove(amount int, from int, to int) {
	for a := 0; a < amount; a++ {
		columns[to] = append([]string{columns[from][0]}, columns[to]...)
		columns[from] = columns[from][1:]
		fmt.Println(columns)
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

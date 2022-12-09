package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

var visited = map[coord]bool{}

var head = coord{0, 0}
var tail = coord{0, 0}

func moveRope(com string) {
	if com == "L" {
		head.x--
	} else if com == "R" {
		head.x++
	} else if com == "U" {
		head.y++
	} else if com == "D" {
		head.y--
	}
}

func tailFollow() {
	if head == tail {
		return
	}
	if head.x-tail.x > 1 {
		tail.y = head.y
		tail.x = head.x - 1
	} else if head.x-tail.x < -1 {
		tail.y = head.y
		tail.x = head.x + 1
	} else if head.y-tail.y > 1 {
		tail.x = head.x
		tail.y = head.y - 1
	} else if head.y-tail.y < -1 {
		tail.x = head.x
		tail.y = head.y + 1
	}
}

func main() {
	// open file
	f, err := os.Open("day9input")

	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		slice := strings.Split(scanner.Text(), " ")
		iterInt, err := strconv.Atoi(slice[1])
		if err != nil {
			fmt.Println("error convert iter var to int")
			panic(err)
		}
		for i := 0; i < iterInt; i++ {
			fmt.Println(slice[0])
			moveRope(slice[0])
			tailFollow()
			visited[tail] = true
		}

	}
	fmt.Println(len(visited))
}

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

var rope = []coord{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}

func moveRope(head coord, com string) coord {
	if com == "L" {
		head.x--
	} else if com == "R" {
		head.x++
	} else if com == "U" {
		head.y++
	} else if com == "D" {
		head.y--
	}
	return head
}

func tailFollow(head coord, tail coord) coord {
	if head == tail {
		return tail
	}
	if head.x-tail.x > 1 && head.y-tail.y > 1 {
		tail.x++
		tail.y++
	} else if head.x-tail.x < -1 && head.y-tail.y < -1 {
		tail.x--
		tail.y--
	} else if head.x-tail.x > 1 && head.y-tail.y < -1 {
		tail.x++
		tail.y--
	} else if head.x-tail.x < -1 && head.y-tail.y > 1 {
		tail.x--
		tail.y++
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

	return tail
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
			rope[0] = moveRope(rope[0], slice[0])
			for i := 1; i < len(rope); i++ {
				fmt.Println(rope)
				rope[i] = tailFollow(rope[i-1], rope[i])
				visited[rope[len(rope)-1]] = true
			}

		}

	}
	fmt.Println(len(visited))
}

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

func getCoords() map[coord]int {
	// open file
	coords := map[coord]int{}
	f, err := os.Open("day8-input")

	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	y := 0

	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), "")

		for x, i := range slice {
			size, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			coords[coord{x, y}] = size
		}

		y++
	}
	return coords
}

func isVisible(c coord, coords map[coord]int) int {
	left := 0
	right := 0
	up := 0
	down := 0
	fmt.Println(c)
	//check left to edge
	for i := c.x - 1; i > -1; i-- {
		if coords[coord{i, c.y}] >= coords[c] {
			fmt.Println("can't see it from the left")
			left++
			break
		} else {
			left++
		}
	}
	//check right to edge
	for i := c.x + 1; i < 99; i++ {
		if coords[coord{i, c.y}] >= coords[c] {
			fmt.Println("can't see it from the right")
			right++
			break
		} else {
			right++
		}
	}
	//check up to edge
	for i := c.y - 1; i > -1; i-- {
		if coords[coord{c.x, i}] >= coords[c] {
			fmt.Println("can't see it from the top")
			up++
			break
		} else {
			up++
		}
	}
	//check down to edge
	for i := c.y + 1; i < 99; i++ {
		if coords[coord{c.x, i}] >= coords[c] {
			fmt.Println("can't see it from the bottom")
			down++
			break
		} else {
			down++
		}
	}
	fmt.Println(left, right, up, down)
	fmt.Println(left * right * up * down)
	return left * right * up * down

}

func main() {
	coords := getCoords()
	count := 0
	for i := range coords {
		c := isVisible(i, coords)
		if c > count {
			count = c
		}
	}
	fmt.Println(len(coords))
	fmt.Println(count)
}

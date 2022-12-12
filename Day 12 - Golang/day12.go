package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord struct {
	x      int
	y      int
	val    string
	parent *coord
}

var alph = makeAlphMap("abcdefghijklmnopqrstuvwxyz")
var coordinates = getCoords()

func pathSearch(start coord, goal coord) []coord {
	open := []coord{start}
	closed := []coord{}

	for len(open) > 0 {
		current := open[0]
		fmt.Println(current)
		fmt.Println(open)

		closed = append(closed, current)

		if current.x == goal.x && current.y == goal.y {
			return path(start, current, closed)
		}

		neighbours := generateNeighbours(current, &current)

		open = append(open, newNeighbours(current, neighbours, open, closed)...)

		open = open[1:]
	}

	fmt.Println("ran out of possible values to traverse")
	return nil
}

func generateNeighbours(current coord, parent *coord) []coord {
	coords := []coord{}
	coords = append(coords, coord{x: current.x + 1, y: current.y, val: getVal(current.x+1, current.y), parent: parent}, coord{x: current.x - 1, y: current.y, val: getVal(current.x-1, current.y), parent: parent}, coord{x: current.x, y: current.y + 1, val: getVal(current.x, current.y+1), parent: parent}, coord{x: current.x, y: current.y - 1, val: getVal(current.x, current.y-1), parent: parent})

	return coords
}

func getVal(x int, y int) string {
	for _, v := range coordinates {
		if v.x == x && v.y == y {
			return v.val
		}
	}
	return ""
}

func newNeighbours(current coord, neighbours []coord, open []coord, closed []coord) []coord {
	newNeighbours := []coord{}

	for _, neighbour := range neighbours {

		if inList(neighbour, open) || inList(neighbour, closed) || !inList(neighbour, coordinates) || !canTraverse(current.val, neighbour.val) {
			fmt.Printf("Current val: %v, with neighbour val: %v, returned false. \n", current.val, neighbour.val)
			continue
		}

		fmt.Printf("Current val: %v, with neighbour val: %v, returned true. \n", current.val, neighbour.val)
		newNeighbours = append(newNeighbours, neighbour)
	}

	return newNeighbours
}

func canTraverse(current string, neighbour string) bool {
	if alph[current]+1 >= alph[neighbour] {
		return true
	} else {
		return false
	}
}

func inList(c coord, cs []coord) bool {
	for _, i := range cs {
		if c.x == i.x && c.y == i.y {
			return true
		}
	}
	return false
}

func path(start coord, goal coord, closed []coord) []coord {
	path := []coord{goal}

	cur := goal

	for cur != start {
		cur = parent(cur, closed)
		fmt.Println(cur)

		path = append([]coord{cur}, path...)
	}

	return path
}

func parent(c coord, closed []coord) coord {
	for _, n := range closed {
		if *c.parent == n {
			return n
		}
	}
	return coord{}
}

func getCoords() []coord {
	coords := []coord{}
	f, err := os.Open("day12test")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	y := 0
	for scanner.Scan() {
		for x, v := range scanner.Text() {
			coords = append(coords, coord{x, y, string(v), nil})
		}
		y++
	}
	return coords
}

func makeAlphMap(s string) map[string]int {
	a := make(map[string]int)

	for i := 0; i < len(s); i++ {
		a[string(s[i])] = i + 1
	}
	return a
}

func main() {
	fmt.Println(alph)
	alph["S"] = 1
	alph["E"] = 26
	fmt.Println(coordinates)

	startVal := coord{}
	goalVal := coord{}

	for _, v := range coordinates {
		if v.val == "S" {
			startVal = coord{v.x, v.y, "a", nil}
		} else if v.val == "E" {
			goalVal = coord{v.x, v.y, "z", nil}
		} else {
			continue
		}
	}

	path := pathSearch(startVal, goalVal)

	fmt.Println(len(path))

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var cpuvar = 1
	var cycle = 1
	var data = []string{}
	datap := 0

	var drawing = [][]string{{"#"}, {}, {}, {}, {}, {}}
	// open file
	f, err := os.Open("day10-input")

	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data = append(data, scanner.Text())

	}
	for i := 0; i < 6; i++ {
		for x := i * 40; cycle < (i+1)*40; x++ {
			if data[datap] == "noop" {
				cycle++
				drawing[i] = append(drawing[i], determineDraw(cpuvar, (cycle-1)-(i*40)))
				fmt.Println(cpuvar, (cycle-1)-(i*40))
				fmt.Println(drawing[i])

			} else {
				slice := strings.Split(data[datap], " ")
				cpuInt, err := strconv.Atoi(slice[1])
				if err != nil {
					panic(err)
				}
				cycle++
				drawing[i] = append(drawing[i], determineDraw(cpuvar, (cycle-1)-(i*40)))
				fmt.Println(cpuvar, (cycle-1)-(i*40))
				fmt.Println(drawing[i])

				cycle++
				cpuvar += cpuInt
				drawing[i] = append(drawing[i], determineDraw(cpuvar, (cycle-1)-(i*40)))
				fmt.Println(cpuvar, (cycle-1)-(i*40))
				fmt.Println(drawing[i])
			}
			datap++
		}

	}
	for _, i := range drawing {
		fmt.Println(strings.Join(i, ""))
	}
}

func determineDraw(cpuval int, pixelPos int) string {
	if cpuval == pixelPos || cpuval+1 == pixelPos || cpuval-1 == pixelPos {
		return "#"
	} else {
		return "."
	}
}

func drawArt(str []string) [][]string {
	result := [][]string{}

	// for i := 40; i < 240; i += 40 {
	// 	newLine := []string{}
	// 	 {
	// 		newLine = append(newLine, str[x])
	// 	}
	// 	result = append(result, newLine)
	// }
	return result
}

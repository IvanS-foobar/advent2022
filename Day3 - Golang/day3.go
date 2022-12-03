package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var alph = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func Intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func main() {
	// open file
	f, err := os.Open("day3-input")

	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	p := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		half1 := strings.Split(string(scanner.Text()[0:(len(scanner.Text())/2)]), "")
		half2 := strings.Split(string(scanner.Text()[(len(scanner.Text())/2):len(scanner.Text())]), "")
		fmt.Println(half1)
		fmt.Println(half2)

		p = p + alph[Intersection(half1, half2)[0]]
		// for i := 0; i < len(scanner.Text()); i++ {
		// 	fmt.Println(string(scanner.Text()[i]))
		// }
	}

	fmt.Println(p)

}

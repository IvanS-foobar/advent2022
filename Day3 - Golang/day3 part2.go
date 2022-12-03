package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func makeAlphMap(s string) map[string]int {
	a := make(map[string]int)
	for i := 0; i < len(s); i++ {
		a[string(s[i])] = i + 1
	}
	return a
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
	alph := makeAlphMap("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	f, err := os.Open("day3-input")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	p := 0
	for scanner.Scan() {
		scan1 := strings.Split(scanner.Text(), "")
		scanner.Scan()

		scan2 := strings.Split(scanner.Text(), "")
		common2 := Intersection(scan1, scan2)
		scanner.Scan()

		scan3 := strings.Split(scanner.Text(), "")

		p = p + alph[Intersection(common2, scan3)[0]]
	}

	fmt.Println(p)

}

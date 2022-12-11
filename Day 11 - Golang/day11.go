package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items       []int
	operation   string
	operationBy int
	divBy       int
	ifTrue      int
	ifFalse     int
	inspected   int
}

func getMonkeys() map[int]*monkey {
	var monkeys = map[int]*monkey{}

	// open file
	f, err := os.Open("day11test")

	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	i := 0
	for scanner.Scan() {

		scanner.Scan()
		items := []int{}
		itemsSplice := strings.Split(scanner.Text(), " ")

		for i := 4; i < len(itemsSplice); i++ {
			itemInt, err := strconv.Atoi(itemsSplice[i][0:2])

			if err != nil {
				panic(err)
			}

			items = append(items, itemInt)
		}

		fmt.Println(itemsSplice[4], items)

		scanner.Scan()
		operationsSplice := strings.Split(scanner.Text(), " ")

		fmt.Println(operationsSplice[6])
		opInt, _ := strconv.Atoi(operationsSplice[7])

		scanner.Scan()
		testSplice := strings.Split(scanner.Text(), " ")
		testInt, _ := strconv.Atoi(testSplice[5])

		fmt.Println(testSplice[5])

		scanner.Scan()
		trueSplice := strings.Split(scanner.Text(), " ")
		trueInt, _ := strconv.Atoi(trueSplice[9])

		fmt.Println(trueSplice[9])

		scanner.Scan()
		falseSplice := strings.Split(scanner.Text(), " ")
		falseInt, _ := strconv.Atoi(falseSplice[9])

		fmt.Println(falseSplice[9])

		monkeys[i] = &monkey{items, operationsSplice[6], opInt, testInt, trueInt, falseInt, 0}

		i++
		scanner.Scan()
	}

	return monkeys
}

func doRound(monkeys map[int]*monkey) map[int]*monkey {
	for i := 0; i < len(monkeys); i++ {
		fmt.Printf("monkey %v \n", i)
		for len(monkeys[i].items) != 0 {
			fmt.Printf("item %v \n", monkeys[i].items[0])
			//monkey inspects item, item is operated based on operation
			if monkeys[i].operation == "*" {
				if monkeys[i].operationBy == 0 {
					monkeys[i].items[0] = monkeys[i].items[0] * monkeys[i].items[0]
				} else {
					monkeys[i].items[0] = monkeys[i].items[0] * monkeys[i].operationBy
				}
			} else if monkeys[i].operation == "+" {
				monkeys[i].items[0] = monkeys[i].items[0] + monkeys[i].operationBy
			}

			fmt.Printf("after first inspection item now %v \n", monkeys[i].items[0])
			monkeys[i].inspected++
			fmt.Printf("monkey %v now inspected %v items \n", i, monkeys[i].inspected)

			//monkey gets bored with item, item gets divided by 3
			monkeys[i].items[0] = monkeys[i].items[0] / 3

			fmt.Printf("after division by 3 item now %v \n", monkeys[i].items[0])

			if monkeys[i].items[0]%monkeys[i].divBy == 0 {
				monkeys[monkeys[i].ifTrue].items = append(monkeys[monkeys[i].ifTrue].items, monkeys[i].items[0])

				fmt.Printf("item is divisible by %v, throwing item to %v \n", monkeys[i].divBy, monkeys[i].ifTrue)

				fmt.Printf("recieivng monkey %v now has these items: %v \n", monkeys[i].ifTrue, monkeys[monkeys[i].ifTrue].items)
			} else {
				monkeys[monkeys[i].ifFalse].items = append(monkeys[monkeys[i].ifFalse].items, monkeys[i].items[0])
				fmt.Printf("item not divisible by %v, throwing item to %v \n", monkeys[i].divBy, monkeys[i].ifFalse)
				fmt.Printf("recieivng monkey %v now has these items: %v \n", monkeys[i].ifFalse, monkeys[monkeys[i].ifFalse].items)
			}

			monkeys[i].items = monkeys[i].items[1:]
			fmt.Printf("removed inspected item, current monkey now has %v items \n", monkeys[i].items)
		}
	}
	return monkeys
}

func main() {
	monkeys := getMonkeys()
	rounds := 20

	for i := 0; i < rounds; i++ {
		fmt.Printf("round %v \n", i+1)
		monkeys = doRound(monkeys)

		for i := range monkeys {
			fmt.Println(monkeys[i])
		}
	}

	inspectedSlice := []int{}
	for _, v := range monkeys {
		inspectedSlice = append(inspectedSlice, v.inspected)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspectedSlice)))

	fmt.Println(inspectedSlice)
	fmt.Println(inspectedSlice[0] * inspectedSlice[1])
}

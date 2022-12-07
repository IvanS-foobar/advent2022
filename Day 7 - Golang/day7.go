package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type dirmap map[string]*dir

type dir struct {
	dirmap dirmap
	file   []int
	parent *dir
	size   int
}

// dirmap rootDirectory:{dirmap a: {}, []int, *dirmap rootDirectory}
// var directories = map[string]dir{}

var rootDirectory = &dir{map[string]*dir{}, []int{}, nil, 0}
var curDir = rootDirectory
var count = 0

func getAnswer(d *dir) int {
	//for each k,v pair in dir.dirmap, if size = 0, call this function on it, else add dir size to total
	for _, v := range d.dirmap {
		if v.size == 0 {
			d.size += getAnswer(v)
		} else {
			d.size += v.size
		}
	}
	//for each file, add file size to total
	for _, v := range d.file {
		d.size += v
	}
	if d.size <= 100000 {
		count += d.size
	}
	//assign size to dir
	return d.size
}

func getAnswer2(d *dir, t float64) float64 {
	current := math.Inf(1)

	if len(d.dirmap) > 0 {
		for _, v := range d.dirmap {
			eval := getAnswer2(v, t)
			if eval >= t && eval < current {
				current = eval
			}
		}

		if float64(d.size) >= t && float64(d.size) < current {
			current = float64(d.size)
		}
	}

	return current

	// current_best := 0
	// for _, v := range d.dirmap {
	// 	if v.size <= t && v.size > current_best {
	// 		current_best = v.size
	// 	}
	// 	if
	// }
	// return current_best
}

func main() {
	f, err := os.Open("day7-input")

	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(curDir)
		fmt.Println(scanner.Text())

		slice := strings.Split(scanner.Text(), " ")

		//if it's a command
		if slice[0] == "$" {
			if slice[1] == "ls" {
				continue
			}
			//changing directories. add a child if does not exist. assign parent as pointer to current directory
			if slice[1] == "cd" && slice[2] != ".." {
				if _, ok := curDir.dirmap[slice[2]]; !ok {
					curDir.dirmap[slice[2]] = &dir{dirmap{}, []int{}, curDir, 0}
					fmt.Println(&curDir)
					curDir = curDir.dirmap[slice[2]]
					fmt.Println("adding dir and pointing to it")
					fmt.Println(curDir.parent)
				} else {
					fmt.Println("dir already exists, pointing to it")
					fmt.Println(curDir.dirmap[slice[2]])
					curDir = curDir.dirmap[slice[2]]
				}

				//point to current dir's parent
			} else if slice[2] == ".." {
				fmt.Printf("moving to %v", *curDir.parent)
				fmt.Println(curDir.parent)
				curDir = curDir.parent
				continue
			}

			//if it's listing a directory
		} else if slice[0] == "dir" {
			//if does not exist, create the directory with the parent directory as current directory
			if _, ok := curDir.dirmap[slice[1]]; !ok {
				fmt.Println("found a newdir, creating")
				curDir.dirmap[slice[1]] = &dir{dirmap{}, []int{}, curDir, 0}
			}
			// file with a data volume of some sort
		} else {
			fmt.Println("adding file to directory")
			i, err := strconv.Atoi(slice[0])
			if err != nil {
				fmt.Println("something went wrong. couldnt convert the file size to int")
			}
			curDir.file = append(curDir.file, i)
		}
	}

	fmt.Println(getAnswer(rootDirectory))
	fmt.Println(count)
	max_space := 70000000
	update_size := 30000000
	available_space := max_space - rootDirectory.size
	target_deletion_size := update_size - available_space
	fmt.Println(target_deletion_size)
	fmt.Println(int(getAnswer2(rootDirectory, float64(target_deletion_size))))
}

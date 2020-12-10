package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	// run w/ -f <filename> to use different test file
	var testFile = flag.String("f", "input.txt", "test data file")
	flag.Parse()
	fmt.Println("testing with:", *testFile)
	file, err := os.Open(*testFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	adapters := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var num int
		fmt.Sscanf(line, "%d", &num)
		adapters = append(adapters, num)
	}
	fmt.Println(adapters)
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	fmt.Println(adapters)
	diff1, diff3 := 0, 0
	for i := 1; i < len(adapters); i++ {
		if adapters[i]-adapters[i-1] == 1 {
			diff1++
		} else if adapters[i]-adapters[i-1] == 3 {
			diff3++
		}
	}
	fmt.Println(diff1, diff3, diff1*diff3)

	/*

		0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, (22)

		from n can go to (working backwards):
		0 -> 1 (8)
		1 -> 4 (8)
		4 -> 5,6,7 (4 + 2 + 2 = 8)
		5 -> 6,7 (2 + 2 = 4)
		6 -> 7 (2)
		7 -> 10 (2)
		10 -> 11, 12 (1 + 1 = 2)
		11 -> 12 (1)
		12 -> 15 (1)
		15 -> 16
		19 -> 22 (1)

		so I could find those working backwards
	*/

	paths := make([]int, len(adapters)) // should maybe use int64 as a precaution but this worked for me
	paths[len(adapters)-1] = 1
	for i := len(adapters) - 1; i >= 0; i-- {
		for j := 1; j <= 3; j++ {
			if i+j < len(adapters) {
				if adapters[i+j]-adapters[i] <= 3 {
					paths[i] += paths[i+j]
				}
			}
		}

	}
	fmt.Println(paths)
	fmt.Println("part2:", paths[0])
	fmt.Println("part1:", diff1*diff3)
}

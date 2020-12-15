package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// run w/ -f <filename> to use different test file
	var testFile = flag.String("f", "input.txt", "test data file")
	flag.Parse()
	fmt.Println("testing with:", *testFile)

	bytes, err := ioutil.ReadFile(*testFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(bytes), ",")

	lastSpoken := make(map[int64]int64)
	turn := int64(1)
	lastNumber := int64(0)
	nextNumber := int64(0)
	for _, line := range lines {
		n := int64(0)
		if _, err := fmt.Sscanf(line, "%d", &n); err != nil {
			fmt.Println(err)
			return
		}
		lastSpoken[n] = turn
		turn++
		lastNumber = n
	}
	fmt.Println(lastSpoken, lastNumber)
	for turn != 30000000 { // it's late so I just changed this from 2020 (part1) to 30000000 (part2) - func it if you want
		spoken := nextNumber
		// fmt.Println("turn", turn, "speaking", nextNumber, "last:", lastSpoken[nextNumber])
		if lastSpoken[nextNumber] == 0 { // hasn't been spoken before
			nextNumber = 0
		} else {
			// fmt.Println("  ", turn, "-", lastSpoken[nextNumber], "=", turn-lastSpoken[nextNumber])
			nextNumber = turn - lastSpoken[nextNumber]
		}
		lastSpoken[spoken] = turn

		turn++
	}
	fmt.Println(nextNumber)

}

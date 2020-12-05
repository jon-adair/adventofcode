package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	highest := 0
	seats := make([]bool, 128*8)
	for scanner.Scan() {
		line := scanner.Text()
		ur := 127
		lr := 0
		uc := 7
		lc := 0

		for i := 0; i < 7; i++ {
			if line[i] == 'F' {
				ur = lr + (ur-lr)/2
			} else {
				lr = lr + (ur-lr+1)/2
			}
		}
		// fmt.Println(lr, ur)
		for i := 7; i < 10; i++ {
			if line[i] == 'L' {
				uc = lc + (uc-lc)/2
			} else {
				lc = lc + (uc-lc+1)/2
			}
		}
		seatID := lr*8 + lc
		// fmt.Println(lc, uc)
		fmt.Println(line, seatID)
		seats[seatID] = true
		if highest < seatID {
			highest = seatID
		}

	}
	myseat := 0
	for i := 1; i < highest; i++ {
		if !seats[i] && seats[i+1] && seats[i-1] {
			myseat = i
		}
	}
	fmt.Println("Part One:", highest)
	fmt.Println("Part Two:", myseat)
}

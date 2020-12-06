package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := int64(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		if x, err := strconv.Atoi(scanner.Text()); err == nil {
				y := (x / 3) - 2
			for y >= 0 {
				total += int64(y)
				fmt.Println(scanner.Text(), y, total)
				y = (y / 3) - 2
			}
		} else {
			fmt.Println(scanner.Text(), err)
		}
	}
	fmt.Println("total:", total)
}

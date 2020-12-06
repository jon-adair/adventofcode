package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// run w/ -test to us short input file
	var testFile = flag.Bool("test", false, "run dev test")
	flag.Parse()
	filename := "input.txt"
	if *testFile {
		filename = "test.txt"
	}
	fmt.Println("testing with:", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0
	totalAll := 0
	people := 0
	ansc := make(map[string]int) // string makes it easier to print debug than byte/rune
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println(ansc, people)
			fmt.Print("done: ")
			for c := 'a'; c <= 'z'; c++ {
				if ansc[string(c)] > 0 {
					fmt.Print(string(c))
				}
				if ansc[string(c)] == people {
					totalAll++
				}
			}
			fmt.Println()
			total += len(ansc)
			fmt.Println("answers:", len(ansc), "total:", total, "totalAll:", totalAll)
			ansc = make(map[string]int)
			people = 0
			fmt.Println()
		} else {
			fmt.Println("line:", line)
			for _, c := range line {
				ansc[string(c)]++
			}
			people++
		}
	}

	// 6947
	// 3398
	fmt.Println("Part One:", total)
	fmt.Println("Part Two:", totalAll)
}

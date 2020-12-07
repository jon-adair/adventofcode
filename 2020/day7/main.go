package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	bags := make(map[string]map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println()
		fmt.Println(line)
		if line != "" {
			l := strings.Split(line, " ")
			fmt.Println(l)
			container := l[0] + " " + l[1]
			bc := make(map[string]int)
			for i := 4; i < len(l); i++ {
				if l[i] == "no" {
					break
				}
				n, err := strconv.Atoi(l[i])
				if err != nil {
					fmt.Println(err)
				}
				cbc := l[i+1] + " " + l[i+2]
				fmt.Println(n, cbc)
				bc[cbc] = n
				i += 3
			}
			bags[container] = bc
			fmt.Println(container)
		}
	}
	fmt.Printf("Parsed: %+v\n", bags)
	total := 0
	for k := range bags {
		fmt.Println("top level:", k, bags[k])
		if k == "shiny gold" {
			continue
		}
		total += contains(k, bags)
	}
	fmt.Println(total)

	fmt.Println("***********")
	total2 := 0
	for k := range bags {
		if k == "shiny gold" {
			fmt.Println("top level:", k, bags[k])
			total2 = count(k, bags)
		}
	}
	fmt.Println("part 1:", total)
	fmt.Println("part 2:", total2-1)
	// 235
	// 158493

}

// does this bc (eventually) contain a shiny gold
// can see that I started with the idea that I'd need to sum the counts expecting that to be part 2 but gave up
func contains(bc string, bags map[string]map[string]int) int {
	fmt.Println("contains:", bc)
	total := 0
	if bc == "shiny gold" {
		fmt.Println("this bag is shiny gold")
		return 1
	}

	fmt.Println("checking contained bags of", bc, ":", bags[bc])
	for k := range bags[bc] {
		fmt.Println("recursing on", k, bags[k])
		if contains(k, bags) > 0 {
			fmt.Println("+++")
			total++
			return 1
		}
	}
	return total
}

// part 2 - count the bags contained by this bag
func count(bc string, bags map[string]map[string]int) int {
	fmt.Println("count:", bc)
	total := 0
	fmt.Println("checking contained bags of", bc, ":", bags[bc])
	for k := range bags[bc] {
		fmt.Println("recursing on", k, bags[bc][k])
		c := count(k, bags)
		fmt.Println("recursed on", k, bags[k], ":", bags[bc][k], c)
		if c == 0 {
			total += bags[bc][k]
		} else {
			total += c * bags[bc][k]
		}
	}
	fmt.Println("each", bc, "contained", total, "other bags")
	return total + 1
}

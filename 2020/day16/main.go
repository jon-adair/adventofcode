package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

type rule struct {
	Name               string
	R1L, R1U, R2L, R2U int
	position           int
}

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
	lines := strings.Split(string(bytes), "\n")

	rules := make([]rule, 0)
	tickets := make([][]int, 0)
	validTickets := make([][]int, 0)
	stage := 0
	for _, line := range lines {
		if line == "" {
			stage++
			continue
		}
		switch stage {
		case 0: // rules
			r := new(rule)
			r.position = -1
			sp := strings.Split(line, ":")
			r.Name = sp[0]
			if _, err := fmt.Sscanf(sp[1], "%d-%d or %d-%d", &r.R1L, &r.R1U, &r.R2L, &r.R2U); err != nil {
				fmt.Println("rule", err)
				return
			}
			rules = append(rules, *r)

		case 1: // our ticket - just put it in with the rest
			if line[:4] == "your" {
				continue
			}
			vals := strings.Split(line, ",")
			ticket := make([]int, 0)
			for _, v := range vals {
				x := 0
				if _, err := fmt.Sscanf(v, "%d", &x); err != nil {
					fmt.Println("our ticket:", err)
					return
				}
				ticket = append(ticket, x)
			}
			tickets = append(tickets, ticket)
		case 2: // nearby tickets
			if line[:4] == "near" {
				continue
			}
			vals := strings.Split(line, ",")
			ticket := make([]int, 0)
			for _, v := range vals {
				x := 0
				if _, err := fmt.Sscanf(v, "%d", &x); err != nil {
					fmt.Println("tickets:", err)
					return
				}
				ticket = append(ticket, x)
			}
			tickets = append(tickets, ticket)
		}
	}
	fmt.Println(rules)
	fmt.Println(tickets)

	sum := 0
	for i := range tickets {
		if i == 0 {
			continue // skip our ticket for now
		}
		validTicket := true
		for _, v := range tickets[i] {
			fmt.Println(v)
			valid := false
			for _, r := range rules {
				if (v >= r.R1L && v <= r.R1U) || (v >= r.R2L && v <= r.R2U) {
					fmt.Println(v, "valid", r)
					valid = true
				}
			}
			if !valid {
				sum += v
				validTicket = false
			}
		}
		if validTicket {
			validTickets = append(validTickets, tickets[i])
		}
	}
	fmt.Println(validTickets)
	fmt.Println("part1", sum)

	/*
		build a map for each column - which rules it fits
		then go through those maps, find one with a single match
		remove that rule from the other maps
		mark that rule's match
		repeat N times

		not a great approach but it works with existing data structs I had
	*/
	fmt.Println("checking")

	matches := make([]map[int]bool, len(validTickets[0]))
	for i := range validTickets[0] {
		matches[i] = make(map[int]bool)

	}

	// see what rules each column matches
	for col := range validTickets[0] {
		for ir, r := range rules {
			valid := true
			for _, vv := range validTickets {
				v := vv[col]
				if !((v >= r.R1L && v <= r.R1U) || (v >= r.R2L && v <= r.R2U)) {
					// fmt.Println(v, col, "invalid")
					valid = false
					break
				}
			}
			if valid {
				matches[col][ir] = true
			}
		}
	}
	fmt.Println(matches)
	// repeatedly find a column that only matches 1 rule
	for range validTickets[0] {
		fmt.Println()
		fmt.Println("pass")
		for col := range validTickets[0] {
			if len(matches[col]) == 1 {
				match := -1
				for ir := range rules {
					if matches[col][ir] {
						match = ir
						break
					}
				}
				fmt.Println("col", col, "has only one match:", matches[col], match)
				rules[match].position = col
				for c := range matches {
					delete(matches[c], match)
				}
			}
			fmt.Println(matches)
		}
	}
	fmt.Println(rules)
	// fmt.Println(validTickets)
	// fmt.Println(tickets)

	prod := 1
	for _, r := range rules {
		if strings.HasPrefix(r.Name, "departure") {
			fmt.Println(r.Name, tickets[0][r.position])
			prod *= tickets[0][r.position]
		}
	}
	fmt.Println("part2:", prod)
	

	fmt.Println("part1", sum)

}

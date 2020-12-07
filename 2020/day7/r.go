package main

import (
	"fmt"
	"regexp"
)

// my original attempt at parsing this with regexp capture groups

func main() {

	line := "light red bags contain 1 bright white bag, 2 muted yellow bags."

	fmt.Println(line)
	// this doesn't work because we only get the last instance of the repeating capture group - lesson learned
	regex := *regexp.MustCompile(`^(\w+ \w+) bags contain (?P<b>[0-9] [a-z]+ [a-z]+ bag[s]?[,.][ ]?)*`)
	// this is ugly but would work for my input
	// regex := *regexp.MustCompile(`^(\w+ \w+) bags contain ([0-9] [a-z]+ [a-z]+ bag[s]?[,.][ ]?)?([0-9] [a-z]+ [a-z]+ bag[s]?[,.][ ]?)?([0-9] [a-z]+ [a-z]+ bag[s]?[,.][ ]?)?([0-9] [a-z]+ [a-z]+ bag[s]?[,.][ ]?)?([0-9] [a-z]+ [a-z]+ bag[s]?[,.][ ]?)?`) // ((\d*) bag[s][,.])*`)

	res := regex.FindStringSubmatch(line)
	// fmt.Printf("%+v\n", res)
	// fmt.Println(len(res))
	for i := range res {
		fmt.Println(res[i])
	}
}

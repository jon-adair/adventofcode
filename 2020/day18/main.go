package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var operands []int
var operators []byte

var part2 bool

func isdigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func state() {
	s := make([]string, len(operators))
	for i := range operators {
		s[i] = string(operators[i])
	}
	fmt.Println("  operands :", operands)
	fmt.Println("  operators:", s)
}

func evaluate(op byte) {
	fmt.Println("  start evaluating", string(op))
	state()
	for {
		if op == 0 { // go until empty
			if len(operators) == 0 {
				return
			}
		} else if op == ')' { // go until we match a paren
			if operators[len(operators)-1] == '(' {
				// done
				operators = operators[:len(operators)-1]
				return
			}
		} else if len(operators) == 0 {
			operators = append(operators, op)
			return
		} else { // compare precedence but part 1 they're the same for + or * (but not '('))
			fmt.Println("  top op:", string(operators[len(operators)-1]))
			if operators[len(operators)-1] == '(' {
				operators = append(operators, op)
				return
			}
			if part2 {
				// check precedence: + comes before * in this case
				// so when we get a *, we need to perform any outstanding +'s
				if op == '+' && operators[len(operators)-1] == '*' {
					operators = append(operators, op)
					return
				}
			}
		}
		x := operands[len(operands)-1]
		operands = operands[:len(operands)-1]
		y := operands[len(operands)-1]
		operands = operands[:len(operands)-1]
		o := operators[len(operators)-1]
		operators = operators[:len(operators)-1]
		z := 0
		if o == '+' {
			z = x + y
			fmt.Println("  performing", x, "+", y)
		} else {
			z = x * y
			fmt.Println("  performing", x, "*", y)
		}
		operands = append(operands, z)
		state()
	}
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

	part2 = true

	/*
		bleh so start over

		2 * 3 + (4 * 5) = 26
		((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 = 13632

		ok so:
		find operand, push it on it's stack
		find operator:
		if it's stack is empty, push to it
		if not, compare precedence to what's on the top of the stack
		while top of stack is >= the new operator, perform that, push result, and pop operator
		if ( push it, if ) keep popping (and performing) until we hit a (
		if at end of string, pop until empty

		so first example
		2 push, * push, 3 push, + compare, perform 2*3, push 6, pop *, push +
		push (, push 4, push *, push 5, hit ), 4*5, push 20 pop *, hit (, hit eol
		perform 6+20, push 26, eol and operand empty

		try 3 + 4 * 5
		push 3, +, 4
		hit * so now go perform 3 + 4 = push 7
		push the *


	*/
	sum := 0
	for _, line := range lines {
		fmt.Println("processing new line:", line)
		operands = make([]int, 0)
		operators = make([]byte, 0)
		p := 0
		for {
			fmt.Println("processing:", line[p:])
			if p == len(line) {
				evaluate(0)
				fmt.Println("result:", operands[0])
				sum += operands[0]
				break
			} else if line[p] == ' ' {
				p++
				continue
			} else if isdigit(line[p]) {
				x := 0
				if n, err := fmt.Sscanf(line[p:], "%d", &x); n == 1 && err == nil {
					fmt.Println("num", x)
					p += len(fmt.Sprintf("%d", x))
					fmt.Println("  now", line[p:])
					operands = append(operands, x)
				}
			} else if line[p] == '(' {
				operators = append(operators, line[p])
				p++
			} else if line[p] == ')' {
				evaluate(line[p])
				// operators = append(operators, line[p])
				p++
			} else if line[p] == '*' || line[p] == '+' {
				if len(operators) == 0 {
					operators = append(operators, line[p])
				} else {
					evaluate(line[p])
				}
				p++
			}
			state()
		}
	}
	fmt.Println("part1:", sum)

	// for _, line := range lines {
	// 	x := 0
	// 	result := 0
	// 	p := 0
	// 	op := '+'
	// 	for {
	// 		fmt.Println("processing:", line[p:])

	// 		if line[p] == ' ' {
	// 			p++
	// 		} else if n, err := fmt.Sscanf(line[p:], "%d", &x); n == 1 && err == nil {
	// 			fmt.Println("num", x)
	// 			p += len(fmt.Sprintf("%d", x))
	// 			fmt.Println("  now", line[p:])
	// 			if op == '+' {
	// 				result += x
	// 			} else {
	// 				result *= x
	// 			}
	// 			fmt.Println("result:", result)
	// 		} else if line[p] == '(' {
	// 			fmt.Println("open paren")
	// 			p++
	// 		} else if line[p] == '+' {
	// 			op = '+'
	// 			fmt.Println("op +")
	// 			p++
	// 		} else if line[p] == '*' {
	// 			op = '*'
	// 			fmt.Println("op *")
	// 			p++
	// 		} else {
	// 			fmt.Println("unk", line[p:])
	// 		}
	// 	}
	// }

}

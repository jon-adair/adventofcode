package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type instr struct {
	opcode   string
	arg      int
	executed bool
}

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

	prog := make([]instr, 0)
	for scanner.Scan() {
		line := scanner.Text()
		ins := new(instr)
		fmt.Sscanf(line, "%s %d", &ins.opcode, &ins.arg)
		prog = append(prog, *ins)
	}
	fmt.Println(prog)
	acc := 0
	pc := 0

	for {
		if prog[pc].executed {
			break
		}
		prog[pc].executed = true
		fmt.Println(pc, prog[pc])
		switch prog[pc].opcode {
		case "nop":
			pc++
		case "acc":
			acc += prog[pc].arg
			pc++
		case "jmp":
			pc = pc + prog[pc].arg
		}
	}
	part1 := acc

	done := false
	for i := range prog {
		acc = 0
		pc = 0
		for j := range prog {
			prog[j].executed = false
		}

		if prog[i].opcode == "nop" {
			prog[i].opcode = "jmp"
		} else if prog[i].opcode == "jmp" {
			prog[i].opcode = "nop"
		} else {
			continue
		}
		fmt.Println("running modified:", prog)
		for {
			if pc >= len(prog) {
				fmt.Println("hit end")
				done = true
				break
			}
			if prog[pc].executed {
				fmt.Println("looped")
				break
			}
			prog[pc].executed = true
			fmt.Println(pc, prog[pc])
			switch prog[pc].opcode {
			case "nop":
				pc++
			case "acc":
				acc += prog[pc].arg
				pc++
			case "jmp":
				pc = pc + prog[pc].arg
			}
		}
		if done {
			break
		}
		if prog[i].opcode == "nop" {
			prog[i].opcode = "jmp"
		} else if prog[i].opcode == "jmp" {
			prog[i].opcode = "nop"
		}

	}
	fmt.Println(part1)

	fmt.Println(acc)
}

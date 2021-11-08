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

	prog := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		for _, x := range s {
			if xx, err := strconv.Atoi(x); err == nil {
				prog = append(prog, xx)
			} else {
				log.Fatal(err)
			}
		}
	}
	fmt.Println(prog)
	temp := make([]int, len(prog))
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			copy(temp, prog)
			execute(temp, n, v)
			if temp[0] == 19690720 {
				fmt.Println(n, v, 100*n+v)
				return
			}
		}
	}

	// fmt.Println(prog)
	// fmt.Println(prog[0])
}

func execute(prog []int, noun, verb int) {
	prog[1] = noun
	prog[2] = verb
	pc := 0
	for {
		switch prog[pc] {
		case 1:
			fmt.Println("1 add", prog[pc+1], prog[pc+2], prog[pc+3])
			prog[prog[pc+3]] = prog[prog[pc+1]] + prog[prog[pc+2]]
		case 2:
			fmt.Println("1 mul", prog[pc+1], prog[pc+2], prog[pc+3])
			prog[prog[pc+3]] = prog[prog[pc+1]] * prog[prog[pc+2]]
		case 99:
			log.Println("99 done")
			return
		default:
			log.Fatal("Unknown opcode", prog[pc])
			return
		}
		fmt.Println(pc, prog)
		pc += 4
	}

}

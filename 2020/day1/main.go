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

	in := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		if x, err := strconv.Atoi(scanner.Text()); err == nil {
			in = append(in, x)
		} else {
			fmt.Println("err parsing int", scanner.Text(), err)
		}
	}
	//fmt.Println(in)

A:
	for i := range in {
		for j := range in {
			if in[i]+in[j] == 2020 {
				fmt.Println(in[i] * in[j])
				break A
			}
		}
	}

B:
	for i := range in {
		for j := range in {
			for k := range in {
				if in[i]+in[j]+in[k] == 2020 {
					fmt.Println(in[i] * in[j] * in[k])
					break B
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

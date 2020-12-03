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

	trees := make([][]bool, 0)

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		treeLine := make([]bool, len(line))
		for i := range line {
			treeLine[i] = (line[i] == '#')
		}
		trees = append(trees, treeLine)
	}

	// for i := range trees {
	// 	for j := range trees[i] {
	// 		if trees[i][j] {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	slopes := [][]int{{3, 1}}
	fmt.Println("A:", hitTheTrees(slopes, trees))

	slopes = [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	fmt.Println("B:", hitTheTrees(slopes, trees))
}

func hitTheTrees(slopes [][]int, trees [][]bool) int {
	width := len(trees[0])
	treeCount := 0
	total := 1
	for s := range slopes {
		treeCount = 0
		y := 0
		x := 0
		for {
			x += slopes[s][0]
			y += slopes[s][1]
			if y >= len(trees) {
				break
			}
			// fmt.Println(x, y, trees[y][x%width])
			if trees[y][x%width] {
				treeCount++
			}
		}
		total *= treeCount
	}
	return total
}

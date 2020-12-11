package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

// would have thrown these on a struct if I'd known I'd add so many

func occupied(row, col int, seats [][]string) bool {
	if row < 0 || row >= len(seats) || col < 0 || col >= len(seats[0]) {
		return false
	}
	return seats[row][col] == "#"
}

func seat(row, col int, seats [][]string) bool {
	if row < 0 || row >= len(seats) || col < 0 || col >= len(seats[0]) {
		return false
	}
	return seats[row][col] != "."
}

func adjocc(row, col int, seats [][]string) (sum int) {
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}
			//fmt.Println(row, col, r, c, occupied(row+r, col+c, seats))
			if occupied(row+r, col+c, seats) {
				sum++
			}
		}
	}
	return
}

// is the first visible seat (if any) in this direction occupied?
func visdir(row, col, rd, cd int, seats [][]string) bool {
	for {
		row += rd
		col += cd
		if row < 0 || row >= len(seats) || col < 0 || col >= len(seats[0]) {
			return false // no seats found
		} else if seat(row, col, seats) {
			return occupied(row, col, seats)
		}
	}
}

// how many visible seats are occupied
func visocc(row, col int, seats [][]string) (sum int) {
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}

			if visdir(row, col, r, c, seats) {
				sum++
			}
		}
	}
	return
}

func dump(seats [][]string) {
	fmt.Println()
	for r := range seats {
		for c := range seats[r] {
			fmt.Print(seats[r][c])
		}
		fmt.Println()
	}

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

	// ugly way to make three copies - one to process, one to store to, and one to reset for part 2
	seats := make([][]string, 0)
	scratch := make([][]string, 0)
	orig := make([][]string, 0)
	for scanner.Scan() {
		mapcol := make([]string, 0)
		scratchcol := make([]string, 0)
		origcol := make([]string, 0)
		line := scanner.Text()
		for i := range line {
			mapcol = append(mapcol, string(line[i]))
			scratchcol = append(scratchcol, string(line[i]))
			origcol = append(origcol, string(line[i]))
		}
		seats = append(seats, mapcol)
		scratch = append(scratch, scratchcol)
		orig = append(orig, origcol)
	}
	dump(seats)

	for {
		changes := 0
		// copy last result
		for r := range seats {
			for c := range seats[r] {
				seats[r][c] = scratch[r][c]
			}
		}

		dump(scratch)
		for r := range seats {
			for c := range seats[r] {
				if seat(r, c, seats) {
					occ := adjocc(r, c, seats)
					// fmt.Println(r, c, seats[r][c], occ)
					if seats[r][c] == "L" && occ == 0 {
						scratch[r][c] = "#"
						changes++
					} else if seats[r][c] == "#" && occ >= 4 {
						scratch[r][c] = "L"
						changes++
					}

				}
			}
		}
		dump(scratch)
		if changes == 0 {
			break
		}
	}
	count1 := 0
	for r := range seats {
		for c := range seats[r] {
			if seats[r][c] == "#" {
				count1++
			}
		}
	}
	fmt.Println(count1)

	//////////////////////
	// reset for part 2
	for r := range seats {
		for c := range seats[r] {
			scratch[r][c] = orig[r][c]
		}
	}

	for {
		changes := 0
		for r := range seats {
			for c := range seats[r] {
				seats[r][c] = scratch[r][c]
			}
		}

		dump(scratch)
		for r := range seats {
			for c := range seats[r] {
				if seat(r, c, seats) {
					occ := visocc(r, c, seats)
					//fmt.Println(r, c, seats[r][c], occ)
					if seats[r][c] == "L" && occ == 0 {
						// fmt.Println("%")
						scratch[r][c] = "#"
						changes++
					} else if seats[r][c] == "#" && occ >= 5 {
						scratch[r][c] = "L"
						changes++
					}

				}
			}
		}
		dump(scratch)
		if changes == 0 {
			break
		}
	}
	count2 := 0
	for r := range seats {
		for c := range seats[r] {
			if seats[r][c] == "#" {
				count2++
			}
		}
	}
	fmt.Println("part1:", count1) // 2261
	fmt.Println("part2:", count2) // 2039

}

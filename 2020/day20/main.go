package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

type tile struct {
	id      int
	pattern [10][10]bool
	matches [4]int // NESW 0123
	used    bool
}

func hMatch(a, b tile) bool {
	for i := 0; i < 10; i++ {
		if a.pattern[9][i] != b.pattern[0][i] {
			return false
		}
	}
	return true
}

func vMatch(a, b tile) bool {
	for i := 0; i < 10; i++ {
		if a.pattern[i][9] != b.pattern[i][0] {
			return false
		}
	}
	return true
}

var vert = false
var fixA = false

func flexMatch(a, b tile) bool {
	if vert {
		return vMatch(a, b)
	} else {
		return hMatch(a, b)
	}

}

func anyMatchRotA(a, b tile) (match bool, aa, bb tile) {
	aa = a
	bb = b
	if fixA {
		return flexMatch(aa, bb), aa, bb
	}
	if match = flexMatch(aa, bb); match {
		return
	}
	aa = rotate90(aa)
	if match = flexMatch(aa, bb); match {
		return
	}
	aa = rotate90(aa)
	if match = flexMatch(aa, bb); match {
		return
	}
	aa = rotate90(aa)
	if match = flexMatch(aa, bb); match {
		return
	}
	aa = rotate90(aa)
	return
}

func anyMatchRotB(a, b tile) (match bool, aa, bb tile) {
	aa = a
	bb = b
	if match, aa, bb = anyMatchRotA(aa, bb); match {
		return
	}
	bb = rotate90(bb)
	if match, aa, bb = anyMatchRotA(aa, bb); match {
		return
	}
	bb = rotate90(bb)
	if match, aa, bb = anyMatchRotA(aa, bb); match {
		return
	}
	bb = rotate90(bb)
	if match, aa, bb = anyMatchRotA(aa, bb); match {
		return
	}
	bb = rotate90(bb)
	return
}

func anyMatchVFlipA(a, b tile) (match bool, aa, bb tile) {
	aa = a
	bb = b
	if fixA {
		return anyMatchRotB(aa, bb)
	}
	if match, aa, bb = anyMatchRotB(aa, bb); match {
		return
	}
	aa = vFlip(aa)
	if match, aa, bb = anyMatchRotB(aa, bb); match {
		return
	}
	aa = vFlip(aa)
	return
}

func anyMatchHFlipA(a, b tile) (match bool, aa, bb tile) {
	aa = a
	bb = b
	if fixA {
		return anyMatchVFlipA(aa, bb)
	}

	if match, aa, bb = anyMatchVFlipA(aa, bb); match {
		return
	}
	aa = hFlip(aa)
	if match, aa, bb = anyMatchVFlipA(aa, bb); match {
		return
	}
	aa = hFlip(aa)
	return
}

func anyMatchVFlipB(a, b tile) (match bool, aa, bb tile) {
	aa = a
	bb = b
	if match, aa, bb = anyMatchHFlipA(aa, bb); match {
		return
	}
	bb = vFlip(bb)
	if match, aa, bb = anyMatchHFlipA(aa, bb); match {
		return
	}
	bb = vFlip(bb)
	return
}

func anyMatchHFlipB(a, b tile) (match bool, aa, bb tile) {
	aa = a
	bb = b
	if match, aa, bb = anyMatchVFlipB(aa, bb); match {
		return
	}
	bb = hFlip(bb)
	if match, aa, bb = anyMatchVFlipB(aa, bb); match {
		return
	}
	bb = hFlip(bb)
	return
}

func anyMatch(a, b tile) (match bool, aa, bb tile) {
	return anyMatchHFlipB(a, b)
}

func rotate90(a tile) (t tile) {
	// to rotate put 0,0 at 9,0
	// 0,1 at 8,0
	// 1,3 at 7,1
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			t.pattern[9-j][i] = a.pattern[i][j]
		}
	}
	t.id = a.id
	return
}

func vFlip(a tile) (t tile) {
	// 0,0 to 0,9
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			t.pattern[i][9-j] = a.pattern[i][j]
		}
	}
	t.id = a.id
	return
}

func hFlip(a tile) (t tile) {
	// 0,0 to 9,0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			t.pattern[9-i][j] = a.pattern[i][j]
		}
	}
	t.id = a.id
	return
}

func dump(t tile) {
	fmt.Println()
	fmt.Println("Tile: ", t.id)
	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			if t.pattern[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

var part2 bool = true

func dumpGrid(tiles [][]tile) {
	for j := 0; j < len(tiles); j++ {
		for i := 0; i < len(tiles); i++ {
			fmt.Printf("%6d", tiles[i][j].id)
		}
		fmt.Println()

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

	i := 0
	tiles := make([]tile, 0)
	for {
		fmt.Println("line:", lines[i])
		if lines[i] == "" {
			i++
			continue
		}
		xx := strings.Split(lines[i], " ")
		t := new(tile)
		fmt.Sscanf(xx[1], "%d", &t.id)
		for y := 0; y < 10; y++ {
			i++
			for x := 0; x < 10; x++ {
				t.pattern[x][y] = lines[i][x] == '#'
			}
		}
		i++

		fmt.Println("adding tile:", t)
		tiles = append(tiles, *t)
		if i == len(lines) {
			break
		}
	}

	/*
		now how to solve it?
		recursively check them?

		oh shit, rotate and flip too


		I guess I need to find unique edges
		should be 4 tiles with 2 of those
		pick one put it in the corner then search out the edges

		can compute a 10-bit number for each edge and it's flip

	*/
	dump(tiles[0])
	dump(rotate90(tiles[0]))
	dump(hFlip(tiles[0]))
	dump(vFlip(tiles[0]))

	fmt.Println(len(tiles))
	var grid [][]tile
	dim := 3
	if len(tiles) == 144 {
		dim = 12
	}
	grid = make([][]tile, dim)
	for i := range grid {
		grid[i] = make([]tile, dim)
	}

	sum := 1
	corner := 0
	for i := range tiles {
		c := 0
		for j := range tiles {
			if i == j {
				continue
			}
			m, _, _ := anyMatch(tiles[i], tiles[j])
			if m {
				// fmt.Println("tile", i, "matches", j)
				c++
			}

		}
		fmt.Println("tile", i, "has", c, "matches")
		if c == 2 {
			fmt.Println("tile", i, "is a corner")
			sum *= tiles[i].id
			switch corner {
			case 0:
				grid[0][0] = tiles[i]
				tiles[i].used = true
			case 1:
				//grid[0][dim-1] = tiles[i]
			case 2:
				//grid[dim-1][0] = tiles[i]
			case 3:
				//grid[dim-1][dim-1] = tiles[i]
			}
			corner++
			// break
		}
	}
	dumpGrid(grid)

	for y := 0; y < dim-1; y++ {
		for x := 0; x < dim-1; x++ {
			// if grid[x][y].id != 0 {
			// 	continue
			// }
			fmt.Println()
			fmt.Println("processing", x, y)

			holes := 0

			// if x > 0 && grid[x-1][y].id == 0 {
			// 	holes++
			// }
			// if y > 0 && grid[x][y-1].id == 0 {
			// 	holes++
			// }
			if x < dim-1 && grid[x+1][y].id == 0 {
				holes++
			}
			if y < dim-1 && grid[x][y+1].id == 0 {
				holes++
			}
			fmt.Println("holes:", holes)

			rotLimit := 1
			if x == 0 && y == 0 {
				rotLimit = 4
			}
			tile1, tile2 := 0, 0
			for rot := 0; rot < rotLimit; rot++ {
				// This is kind of a mess.
				// Trying to meet two needs which is making it too complex
				// I have a corner tile in 0,0 and need to find the right orientation for it
				// so I need to rotate and flip (though don't need to) it until I have a right and bottom match
				// Then once I have those 3 tiles in place, oriented properly, I just go through the rest of the grid
				// at each grid square, find a right and bottom match (unless I'm at the right / bottom edge), orient it, and store it
				found := 0
				if rotLimit != 1 {
					grid[x][y] = rotate90(grid[x][y])
				}
				fixA = true
				if x < dim-2 {
					vert = false
					for i := range tiles {
						if grid[x][y].id == tiles[i].id || tiles[i].used {
							continue
						}
						m, a, b := anyMatch(grid[x][y], tiles[i])
						if m {
							fmt.Println(rot, "match")
							// dump(a)
							// dump(b)
							found++
							grid[x][y] = a
							grid[x+1][y] = b
							tile1 = i
						}
					}
				}

				if y < dim-2 {
					vert = true
					for i := range tiles {
						if grid[x][y].id == tiles[i].id || tiles[i].used {
							continue
						}
						m, a, b := anyMatch(grid[x][y], tiles[i])
						if m {
							fmt.Println(rot, "match")
							dump(a)
							dump(b)
							found++
							grid[x][y] = a
							grid[x][y+1] = b
							tile2 = i
						}
					}
				}
				if found == holes {
					fmt.Println("right orientation is", rot)
					tiles[tile1].used = true
					tiles[tile2].used = true
					break
				}
			}
			dumpGrid(grid)
		}
	}
	for i := range tiles {
		if !tiles[i].used {
			fmt.Println(tiles[i].id)
		}
	}

	fmt.Printf("1951    2311    3079\n2729    1427    2473\n2971    1489    1171\n")

	// ok now I have a corner tile in the corner but unsure of the right rotation / flip

	/*
		// ok so let's be dumb and just look for a corner
		for i := range tiles {
			// try each tile as top-left corner

			// so I need to check each edge of the tile against every other edge
			for rot := 0; rot < 4; rot++ { // check each rotation
				tiles[i] = rotate90(tiles[i])

			}
		}

		// nah start by finding out how many matches there are for each edge of each tile
		for i := range tiles {
			// so I need to check each edge of the tile against every other edge
			for rot := 0; rot < 4; rot++ { // check each rotation
				t := tiles[i]
				for r := 0; r < rot; r++ {
					t = rotate90(t)
				}
				for j := range tiles {
					if i == j {
						continue
					}
					if hMatch(t, tiles[j]) {

							// but I need to rotate and flip everything else
							// this is really probably better served with some sort of map of ints representing each edge
							// so run through each tile and compute the int for each edge and the reverse
							// build a map that counts those
							// find ...

							// you know I think it's only 12x12
							// too big to just exhaustively search? or randomly fill and verify?

							// been thinking of starting in the center and recursively grow out from there
							// really no different than starting on edge

							// so instead what if I start with each tile as a corner and grow from there?


					}
				}

			}
		}
	*/

	fmt.Println("sum:", sum)

}

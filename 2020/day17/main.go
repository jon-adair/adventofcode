package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

// I again trashed my part 1 so this is part 2 only

// I had a decent structure so adding w was pretty easy but if I did it again knowing part 2,
// I'd maybe use a slice inside coords but that would make some of the funcs less clear shrug

type coords struct {
	x, y, z, w int
}

type grid map[coords]bool

func (g grid) getrange() (min, max coords) {
	for k := range g {
		if k.x > max.x {
			max.x = k.x
		}
		if k.y > max.y {
			max.y = k.y
		}
		if k.z > max.z {
			max.z = k.z
		}
		if k.w > max.w {
			max.w = k.w
		}

		if k.x < min.x {
			min.x = k.x
		}
		if k.y < min.y {
			min.y = k.y
		}
		if k.z < min.z {
			min.z = k.z
		}
		if k.w < min.w {
			min.w = k.w
		}

	}
	return
}

func (g grid) dump() {
	min, max := g.getrange()
	for w := min.w; w <= max.w; w++ {
		fmt.Printf("w=%d\n", w)
		for z := min.z; z <= max.z; z++ {
			fmt.Printf("z=%d, (%d,%d)\n", z, min.x, min.y)
			for y := min.y; y <= max.y; y++ {
				for x := min.x; x <= max.x; x++ {
					if g[coords{x, y, z, w}] {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println()
			}
		}
	}
	fmt.Println(min, max)
	// fmt.Println(g)
}

func (g grid) adj(x, y, z, w int) (sum int) {
	for xx := -1; xx <= 1; xx++ {
		for yy := -1; yy <= 1; yy++ {
			for zz := -1; zz <= 1; zz++ {
				for ww := -1; ww <= 1; ww++ {
					if xx == 0 && yy == 0 && zz == 0 && ww == 0 {
						continue
					}
					if g[coords{x + xx, y + yy, z + zz, w + ww}] {
						sum++
					}
				}
			}
		}
	}
	return
}

func (g grid) grow() grid {
	min, max := g.getrange()
	result := make(map[coords]bool)
	for k, v := range g {
		result[k] = v
	}
	for ww := min.w - 1; ww <= max.w+1; ww++ {
		for zz := min.z - 1; zz <= max.z+1; zz++ {
			for yy := min.y - 1; yy <= max.y+1; yy++ {
				for xx := min.x - 1; xx <= max.x+1; xx++ {
					a := g.adj(xx, yy, zz, ww)
					// fmt.Println("adj", xx, yy, zz, "=", a)
					if g[coords{xx, yy, zz, ww}] {
						if a == 2 || a == 3 {
							// fmt.Println("retain", xx, yy, zz)
							result[coords{xx, yy, zz, ww}] = true
						} else {
							// fmt.Println("delete", xx, yy, zz)
							delete(result, coords{xx, yy, zz, ww})
						}
					} else if !g[coords{xx, yy, zz, ww}] && a == 3 {
						// fmt.Println("add", xx, yy, zz)
						result[coords{xx, yy, zz, ww}] = true
					}
				}
			}
		}
	}
	return result
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

	var active grid
	active = make(map[coords]bool)
	// var max, min coords

	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				active[coords{x, y, 0, 0}] = true
			}
		}
	}
	// .#.
	// ..#
	// ###
	active.dump()
	// fmt.Println(active.adj(2, 2, 1))
	for i := 1; i <= 6; i++ {
		active = active.grow()
	}
	// fmt.Println("grow 1")
	// active = active.grow()
	active.dump()
	fmt.Println(len(active))
}

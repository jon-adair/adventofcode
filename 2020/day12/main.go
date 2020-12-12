package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func rotate(x, y, deg int) (nx, ny int) {
	switch deg {
	case 0:
		nx = x
		ny = y
	case 90:
		nx = y
		ny = -x
	case 180:
		nx = -x
		ny = -y
	case 270:
		nx = -y
		ny = x
	}
	return
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

	x, y := 0, 0
	dir := 90
	for _, line := range lines {
		cmdc := ' '
		param := 0
		if n, err := fmt.Sscanf(line, "%c%d", &cmdc, &param); n != 2 || err != nil {
			fmt.Println(err)
		}
		cmd := string(cmdc)
		// fmt.Println(cmd, param)
		switch cmd {
		case "L":
			dir = (dir - param) % 360
			if dir < 0 {
				dir += 360
			}
		case "R":
			dir = (dir + param) % 360
		case "F":
			switch dir {
			case 0:
				y += param
			case 180:
				y -= param
			case 90:
				x += param
			case 270:
				x -= param
			}
		case "N":
			y += param
		case "S":
			y -= param
		case "E":
			x += param
		case "W":
			x -= param

		}
		fmt.Printf("%s %3d -> (%4d, %4d) %3d\n", cmd, param, x, y, dir)
	}
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	part1 := fmt.Sprint("part 1: ", x, y, x+y)
	fmt.Println("part1", x, y, x+y)

	x, y = 0, 0
	wx, wy := 10, 1
	for _, line := range lines {
		cmdc := ' '
		param := 0
		if n, err := fmt.Sscanf(line, "%c%d", &cmdc, &param); n != 2 || err != nil {
			fmt.Println(err)
		}
		cmd := string(cmdc)
		// fmt.Println(cmd, param)
		switch cmd {
		case "L":
			p := 360 - param // turns out, left is different than right 🤷‍♀️
			fmt.Print("rotate ", p, wx, wy, " -> ")
			wx, wy = rotate(wx, wy, p)
			fmt.Println(wx, wy)
		case "R":
			fmt.Print("rotate ", param, wx, wy, " -> ")
			wx, wy = rotate(wx, wy, param)
			fmt.Println(wx, wy)
		case "F":
			x += wx * param
			y += wy * param
		case "N":
			wy += param
		case "S":
			wy -= param
		case "E":
			wx += param
		case "W":
			wx -= param
		}
		fmt.Printf("%s %3d -> (%4d, %4d) (%4d %4d)\n", cmd, param, x, y, wx, wy)
	}
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	fmt.Println(part1)                // 938
	fmt.Println("part 2:", x, y, x+y) // 54404

}

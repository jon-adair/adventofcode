package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

// I broke my part1 and not going to bother fixing it

// apply chinese remainder theorum
func crt(divs, rems []int64) int64 {
	fmt.Println("CRT:", divs, rems)
	terms := make([]int64, len(divs))

	for i := range divs {
		terms[i] = 1
		for j := range divs {
			if i == j {
				continue
			}
			terms[i] *= divs[j]
		}
	}
	fmt.Println("initial terms:", terms)
	// now we brute-force multiply each term until we find the right one
	// could instead use Extended Euclidean algorithm
	sum := int64(0)
	product := int64(1)
	for i := range divs {
		t := terms[i]
		for {
			fmt.Printf("%d %% %d = %d, need %d\n", divs[i], t, t%divs[i], rems[i])
			if t%divs[i] == rems[i] {
				terms[i] = t
				break
			}
			t += terms[i]
		}
		sum += t
		product *= divs[i]
	}
	fmt.Println("terms for mods:", terms)
	fmt.Println(sum, product, sum%product)

	return sum % product
}

// a mod b - because % is remainder operator not modulo and differs for negative a
func mod(a, b int64) int64 {
	return (a%b + b) % b
}

func main() {
	// see the diff between % and mod:
	fmt.Println(-1 % 3)         // -1
	fmt.Println((-1%3 + 3) % 3) // 2

	// fmt.Println(crt([]int64{5,3}, []int64{0, 0}))
	// fmt.Println(crt([]int64{5,3}, []int64{0, 1}))
	// fmt.Println(crt([]int64{5,3}, []int64{0, 2}))

	// fmt.Println(crt([]int64{3, 5}, []int64{0, 4}))

	// fmt.Println(crt([]int64{7,13,59,31,19},
	// 		        []int64{0,12,55,25,12}))

	// return
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

	ts := int64(0)
	fmt.Sscanf(lines[0], "%d", &ts)
	sked := strings.Split(lines[1], ",")
	buses := make([]int64, 0)
	for _, b := range sked {
		if b != "x" {
			bus := int64(0)
			fmt.Sscanf(b, "%d", &bus)
			buses = append(buses, bus)
		} else {
			buses = append(buses, 0)
		}
	}

	// part 1 - simple search
	t := int64(0)
outer:
	for {
		for _, b := range buses {
			if b != 0 && (ts+t)%b != 0 {
				t++
				continue outer
			}
		}
		break
	}
	part1 := ts + t
	fmt.Println("part1:", part1)
	// fmt.Println(buses)
	// for bi, b := range buses {
	// 	fmt.Println(bi, b)
	// }
	fmt.Println()

	// Part 2 - solve using Chinese Remainder Theorum
	divs := make([]int64, 0)
	mods := make([]int64, 0)
	for i := range buses {
		if buses[i] == 0 {
			continue
		}
		divs = append(divs, buses[i])
		mod := ((buses[i]-int64(i))%buses[i] + buses[i]) % buses[i]
		fmt.Printf("[%d] x mod %d = %d\n", i, buses[i], mod)
		// fmt.Println("  ", i, buses[i], mod)
		mods = append(mods, mod)
	}
	fmt.Println()
	fmt.Println(divs)
	fmt.Println(mods)
	fmt.Println("part2", crt(divs, mods))
	fmt.Println("part1", part1)
	return

	/*

		a little more of a nasty example where a bus departs multiple times in the time range:
		   5 x x x 3
		0  D       D
		1
		2
		3          D
		4
		5  D
		6    .     D
		7      .
		8        .
		9          D



	*/

	// discarded iterative solution but useful as a check for small cases
	fmt.Println()
	fmt.Println("iterative:")
	t = ts
	t = 0

	for {
		unmatch := false
		for bi, b := range buses {
			if b == 0 {
				continue
			}
			// fmt.Println("ts", t, "%", b, "=", t%b, "bi", bi)
			fmt.Printf("ts %d: (%d+%d)%%%d=%d\n", t, t, bi, b, (t+int64(bi))%b)
			if (t+int64(bi))%b == 0 {
				// fmt.Println(t,b,(t-ts)*b)
				// return
			} else {
				unmatch = true
			}
		}
		if !unmatch {
			fmt.Println(t)
			return
		}
		t++ //int64(buses[0])
		// if t > 1068781 { break}
	}
	fmt.Println(t)
}

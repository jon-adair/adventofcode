package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

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

	var maska, masko int64

	storage := make(map[int64]int64)
	maskstr := ""

	for _, line := range lines {
		if line[:4] == "mask" {
			fmt.Println()
			fmt.Println("new mask:")
			maska, masko = 0, 0
			maskstr = line[7:]
			fmt.Println(maskstr, " new mask")
			for i := range maskstr {
				maska = maska << 1
				masko = masko << 1
				if maskstr[i] == '0' {
					maska++
				} else if maskstr[i] == '1' {
					masko++
				}
				maskstr = strings.Replace(maskstr, "X", ".", -1)
			}
			// maska = !maska
			fmt.Println(maska, masko)
			fmt.Printf("%s string\n%036b maska\n%036b masko\n", maskstr, maska, masko)
			continue
		}

		a, v := int64(0), int64(0)
		if n, err := fmt.Sscanf(line, "mem[%d] = %d", &a, &v); n != 2 || err != nil {
			fmt.Println("error parsing", err)
			return
		}

		// maska is misnamed - not really an AND mask, but a clear (AND NOT) mask
		storage[a] = (v | masko) &^ maska
		fmt.Printf("\nwrite [%d] = %d\n%036b v\n%s mask\n%036b result\n", a, v, v, maskstr, storage[a])

	}
	sum := int64(0)
	for _, s := range storage {
		sum += s
	}
	fmt.Println("part1:", sum)

	///// part 2

	storage = make(map[int64]int64)
	maskx := int64(0)
	for _, line := range lines {
		if line[:4] == "mask" {
			fmt.Println()
			fmt.Println("new mask:")
			maska, masko, maskx = 0, 0, 0
			maskstr = line[7:]
			fmt.Println(maskstr, " new mask")
			for i := range maskstr {
				maska = maska << 1
				masko = masko << 1
				maskx = maskx << 1
				if maskstr[i] == '1' {
					masko++
				} else if maskstr[i] == 'X' {
					maskx++
				}
				//maskstr = strings.Replace(maskstr, "1", ".", -1)
			}
			// maska = !maska
			fmt.Println(maska, masko, maskx)
			fmt.Printf("%s string\n%036b maska\n%036b masko\n%036b maskx\n", maskstr, maska, masko, maskx)
			continue
		}

		a, v := int64(0), int64(0)
		if n, err := fmt.Sscanf(line, "mem[%d] = %d", &a, &v); n != 2 || err != nil {
			fmt.Println("error parsing", err)
			return
		}

		addr := (a | masko) &^ maskx
		fmt.Printf("\nwrite [%d] = %d\n%036b addr\n%s mask\n%036b init result\n", a, v, a, maskstr, addr)
		// got the original unfuzzed addr, now we'll resolve all the X bits
		// yeah this should be done once per mask but it's late 😴
		t := maskx
		p := int64(1)
		addrng := make([]int64, 1)
		for t > 0 {
			if t%2 == 1 {
				for _, r := range addrng {
					addrng = append(addrng, r+p)
				}
			}
			t = t >> 1
			p = p << 1
		}
		// addrng now has a list of all the offsets to write to
		fmt.Println("writing to deltas (from", addr, "):", addrng)
		for _, r := range addrng {
			fmt.Println("writing to ", addr+r)
			storage[addr+r] = v
		}

	}
	fmt.Println("part 1:", sum) // 10452688630537
	sum = int64(0)
	for _, s := range storage {
		sum += s
	}
	fmt.Println("part 2:", sum) //2881082759597

}

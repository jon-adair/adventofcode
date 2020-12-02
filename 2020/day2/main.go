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

	validA := 0
	validB := 0
	for scanner.Scan() {
		counts := make(map[byte]int)
		hi := 0
		lo := 0
		c := byte('.')
		pw := ""
		fmt.Println(scanner.Text())
		if n, err := fmt.Sscanf(scanner.Text(),"%d-%d %c: %s",&lo,&hi,&c,&pw); n==4 && err == nil {
			//fmt.Println(hi,lo,c,pw)
			for i := range pw {
				counts[pw[i]]++
			}
			//fmt.Println(counts)
			//fmt.Println(counts[c])
			if counts[c] >= lo && counts[c] <= hi {
				validA++
			}
			if (pw[lo-1] == c) != (pw[hi-1] == c) {
				validB++
			}
		} else {
			fmt.Println("err parsing", scanner.Text(), err)
		}
	}
	fmt.Println("A:", validA)
	fmt.Println("B:", validB)
}

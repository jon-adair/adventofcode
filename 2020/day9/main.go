package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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
	pre := 25
	nums := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var num int
		fmt.Sscanf(line, "%d", &num)
		nums = append(nums, num)
	}
	invalidNum := -1
invalid:
	for i := pre; i < len(nums); i++ {
		foundValid := false
	valid:
		for j := 1; j <= pre; j++ {
			for k := 1; k <= pre; k++ {
				if j == k {
					continue
				}
				if nums[i] == nums[i-j]+nums[i-k] {
					fmt.Println("valid", i, j, k, nums[i], nums[j], nums[k])
					foundValid = true
					break valid
				}
			}
		}
		if !foundValid {
			fmt.Println("invalid", nums[i], i)
			invalidNum = nums[i]
			break invalid
		}
	}

	for i := 0; i < len(nums); i++ {
		sum := 0
		min, max := nums[i], nums[i]
		fmt.Println(i, "summing:")
		for j := 0; j < len(nums)-i; j++ {
			fmt.Println(nums[i+j], min, max, "min max")
			if min > nums[i+j] {
				min = nums[i+j]
			}
			if max < nums[i+j] {
				max = nums[i+j]
			}

			fmt.Println(sum, "+", nums[i+j], i, j)
			sum += nums[i+j]
			// fmt.Println(i, j, sum)
			if sum == invalidNum {
				fmt.Println("found", i, j, nums[i]+nums[j], min, max, min+max)

				return
			}
		}
		fmt.Println("sum:", sum)
	}

}

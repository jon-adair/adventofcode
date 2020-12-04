package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	//"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0
	passport := ""
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		passport = passport + line + " "
		if line == "" {
			if checkPassport(passport, true) {
				fmt.Println("valid")
				total++
			} else {
				fmt.Println("invalid")
			}
			passport = ""
			fmt.Println()

		}
	}
	fmt.Println(total)

}

func checkPassport(p string, validate bool) bool {
	fmt.Println(p)
	pp := strings.Split(p, " ")
	// fmt.Println(pp)
	need := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	m := make(map[string]bool) // should maybe have counted instead to ensure 1 and only 1 of each
	for _, ppp := range pp {
		pppp := strings.Split(ppp, ":")
		m[pppp[0]] = true
		if validate {
			fmt.Println("validating", pppp)
			switch pppp[0] {
			case "byr":
				if len(pppp[1]) != 4 {
					fmt.Println("wrong len")
					return false
				}
				if v, err := strconv.Atoi(pppp[1]); err == nil {
					if v < 1920 || v > 2002 {
						fmt.Println("bad range", v)
						return false
					}
				} else {
					fmt.Println(err)
					return false
				}
			case "iyr":
				if len(pppp[1]) != 4 {
					fmt.Println("wrong len")
					return false
				}
				if v, err := strconv.Atoi(pppp[1]); err == nil {
					if v < 2010 || v > 2020 {
						fmt.Println("bad range", v)
						return false
					}
				} else {
					fmt.Println(err)
					return false
				}
			case "eyr":
				if len(pppp[1]) != 4 {
					fmt.Println("wrong len")
					return false
				}
				if v, err := strconv.Atoi(pppp[1]); err == nil {
					if v < 2020 || v > 2030 {
						fmt.Println("bad range", v)
						return false
					}
				} else {
					fmt.Println(err)
					return false
				}
			case "hgt":
				if !strings.HasSuffix(pppp[1], "cm") && !strings.HasSuffix(pppp[1], "in") {
					return false
				}
				if v, err := strconv.Atoi(pppp[1][:len(pppp[1])-2]); err == nil {
					if strings.HasSuffix(pppp[1], "cm") && (v < 150 || v > 193) {
						fmt.Println("bad range", v)
						return false
					}
					if strings.HasSuffix(pppp[1], "in") && (v < 59 || v > 76) {
						fmt.Println("bad range", v)
						return false
					}
				} else {
					fmt.Println(err)
					return false
				}
			case "hcl":
				if matched, err := regexp.MatchString("^#[0-9a-f]{6}$", pppp[1]); err != nil || !matched {
					fmt.Println("bad hair color", pppp[1], err)
					return false
				}

			case "ecl":
				if matched, err := regexp.MatchString("^amb|blu|brn|gry|grn|hzl|oth$", pppp[1]); err != nil || !matched {
					fmt.Println("bad ecl", pppp[1], err)
					return false
				}
			case "pid":
				if matched, err := regexp.MatchString("^\\d{9}$", pppp[1]); err != nil || !matched {
					fmt.Println("bad pid", pppp[1], err)
					return false
				}
			}
		}
	}
	fmt.Println(m)
	for _, n := range need {
		y, _ := m[n]
		if !y {
			fmt.Println("missing", n)
			return false
		}
	}
	return true
}

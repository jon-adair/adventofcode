package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type node struct {
	rule  int
	leaf  string
	nodes [][]int
}

var part2 bool = true

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

	nodes := make(map[int]node)
	li := 0
	for _, line := range lines {
		if line == "" {
			li++
			break
		}
		fmt.Println("line:", line)
		s := strings.Split(line, ":")
		newNode := new(node)
		// rule := 0
		fmt.Sscanf(s[0], "%d", &newNode.rule)
		if strings.Contains(s[1], "\"") {
			newNode.leaf = string(s[1][2])
		} else {
			newNode.nodes = make([][]int, 0)
			ss := strings.Split(s[1], "|")
			fmt.Println("  ss:", ss)
			for i, sss := range ss {
				fmt.Println("parsing:", sss)
				n := make([]int, 3) // I know I have at most 3 in my input
				nn, err := fmt.Sscanf(sss, "%d %d %d", &n[0], &n[1], &n[2])
				if err != nil && err.Error() != "EOF" {
					fmt.Println("err parsing", err, nn)
					return
				}
				n = n[:nn]
				fmt.Println("read", nn, "ints:", n)
				newNode.nodes = append(newNode.nodes, make([]int, 0))
				for _, x := range n {
					newNode.nodes[i] = append(newNode.nodes[i], x)
				}
			}
		}
		nodes[newNode.rule] = *newNode
		li++
	}
	fmt.Println(nodes)

	// now I could build a parser - and started to - but what if I instead generate a regexp to handle it?
	// worked great for part 1 and part 2 hurt a little but not too bad.
	// for the curious, my final 23661 character regexp (with 6 pairs on node11) is in regexp.txt

	exp := make(map[int]string)

	//  8: 42 | 42 8
	// 11: 42 31 | 42 11 31

	for {
		if exp[0] != "" {
			break
		}
		for i, node := range nodes {
			if exp[i] != "" {
				continue
			}
			if node.leaf != "" {
				exp[i] = node.leaf
			} else {
				fmt.Println("node:", node)
				fmt.Println("  ", len(node.nodes))
				allFound := true
				expr := ""
				if node.rule == 8 {
					fmt.Println("8[0]:", exp[node.nodes[0][0]])
					// fmt.Println("8[1]:", node.nodes[1])
					// fmt.Println("8[2]:", node.nodes[2])
				}

				for _, alt := range node.nodes {
					for _, c := range alt {
						if exp[c] != "" {
							expr += exp[c]
						} else {
							allFound = false
						}
					}
					expr += "|"
				}
				expr = expr[:len(expr)-1]
				expr = "(" + expr + ")"
				if allFound {
					if part2 {
						// hack up the expression for the two special cases
						if node.rule == 8 {
							// pretty easy, just allow 1+ of this term
							//  8: 42 | 42 8

							expr = expr + "+"
						}
						if node.rule == 11 {
							// a real pain
							// was 11: 42 31
							// now 11: 42 31 | 42 11 31
							// so I'm going to be a horrible hack and just insert some number of terms
							// by hand to match for my data. 6 was enough for me. 20 is safer but sloow
							fmt.Println()
							fmt.Println("node 11:", expr)
							expr = ""
							for i := 1; i < 6; i++ {
								expr += "("
								for j := 0; j < i; j++ {
									expr += exp[42]
								}
								for j := 0; j < i; j++ {
									expr += exp[31]
								}
								expr += ")|"
							}
							expr = expr[:len(expr)-1]
							expr = "(" + expr + ")"
							fmt.Println("node 11:", expr)
							fmt.Println()
						}
					}
					exp[i] = expr
				}
			}

		}
		fmt.Println("exp:", exp)
	}
	matcher := "^" + exp[0] + "$"
	fmt.Println("exp:", exp)
	fmt.Println("regexp:", matcher)

	sum := 0
	for _, line := range lines {
		if li > 0 {
			li--
			continue
		}
		//fmt.Println("checking:", line)
		matched, err := regexp.MatchString(matcher, line)
		if err != nil {
			fmt.Println("error matching:", err)
			return
		}
		if matched {
			fmt.Println(" matched:", line)
			sum++
		}
	}
	fmt.Println("sum:", sum)

}

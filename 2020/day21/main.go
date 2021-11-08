package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

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

	// mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
	// trh fvjkl sbzzf mxmxvkd (contains dairy)
	// sqjhc fvjkl (contains soy)
	// sqjhc mxmxvkd sbzzf (contains fish)

	// so follow this logic
	// store each allergen and the list of possible ingredients it could be
	// then look at a union of those lists and see if it's unique
	// then once I know one, can remove that from other lists
	// so start with dairy. The only intersection is mxmxvkd so that's dairy
	// then fish the only intersection (once mxm.. is struck) is sq..
	// then soy is fvj, leaving trh as the only soy

	// what I need then is for each allergen, an array of sets: map[string]bool
	// also slice or rather a map of allergen names -> ingredients
	// remember part1 just needs to know how many times non-allergens appear
	// so I need to keep a list of ingredients for each item - a map[string]bool
	// well the ultimate part1 answer is counting what's left after all the allergens are struck so I'd have that?

	// ok so end result needs to be a list of ingredients that aren't allergens
	// and then how many times they appear - can re-count original or

	ingredientList := make([]map[string]bool, 0)
	allergenList := make([]map[string]bool, 0)
	// ingredientList := make([]map[string]bool, 0)
	// allergenList := make([]map[string]bool, 0)
	for _, line := range lines {
		s := strings.Split(line, " (contains ")
		// fmt.Println(s)
		// fmt.Println(s[0])
		// fmt.Println(s[1])
		ingredients := strings.Split(s[0], " ")
		allergens := strings.Split(s[1][:len(s[1])-1], ", ")
		// fmt.Println(ingredients)
		// fmt.Println(allergens)
		// fmt.Println()
		ingr := make(map[string]bool)
		for _, i := range ingredients {
			ingr[i] = true
		}
		ingredientList = append(ingredientList, ingr)
		alleg := make(map[string]bool)
		allegP := make(map[string]bool)

		for _, i := range allergens {
			alleg[i] = true
			allegP[i] = true
		}
		allergenList = append(allergenList, alleg)

	}
	fmt.Println(allergenList)
	fmt.Println(ingredientList)

	// allergenPossibles := make(map[string][]map[string]bool, 0)

}

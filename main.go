package main

import "fmt"

func main() {
	recipe := ingredients{
		spinachRaw().g(111),
		hempSeeds().g(27),
		miniPeppersEach().servings(3),
		eggEach().servings(3),
		coconutOil().g(20),
		ranch().g(37),
	}.sortBy(byDollars)
	fmt.Print(header, breaks, recipe, breaks, recipe.sum("TOTAL"))
}

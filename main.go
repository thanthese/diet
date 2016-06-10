package main

import "fmt"

func main() {
	recipe := ingredients{
		spinachRaw().g(111),
		hempSeeds().g(27),
		miniPeppersEach().count(3),
		eggEach().count(3),
		coconutOil().g(20),
		ranch().g(37),

		spinachRaw().g(103),
		hempSeeds().g(24),
		coconutOil().g(15),
		cucumber().g(83),
		herring().count(1),
		eggEach().count(2),
		ranch().g(23),
	}.sort(byProtein)
	fmt.Print(header, breaks, recipe, breaks, recipe.sum("TOTAL"))
}

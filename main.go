package main

import "fmt"

func main() {
	recipe := ingredients{
		eggEach().count(1),
		spinachRaw().g(85),
		sardinesCan().count(1),
		chiaSeeds().g(10),
		italian().g(29),

		spinachRaw().g(81),
		avocado().g(151),
		chiaSeeds().g(10),
		cucumber().g(82),
		italian().g(31),
		hempSeeds().g(12),
		eggEach().count(1),
	}.sort(byCalories)
	fmt.Print(header, breaks, recipe, breaks, recipe.sum("TOTAL"))
}

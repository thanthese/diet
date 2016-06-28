package main

import "fmt"

func main() {
	recipe := ingredients{
		eggEach(3),
		sardinesCanBonesEach(1),
		onion(172 / 2),
		redPepper(172 / 2),
		cheese(202 - 172),
		hempSeeds(16),
		oliveOil(11),
	}.sort(byCalories)
	fmt.Print(header, breaks, recipe, breaks, recipe.sum("TOTAL"))
}

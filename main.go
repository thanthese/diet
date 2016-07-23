package main

import "fmt"

func main() {
	recipe := ingredients{
		/*
			traditional smoothie
				almonds
				hemp seeds

				lentils
				spinach

				apples
				strawberries
		*/

		// smoothie
		almonds(100),
		lentils(400),
		strawberries(200),

		// making up the difference
		sardinesCanBonesEach(1),
		eggEach(3),
		oliveOil(40),
	}.sort(byCalories)
	fmt.Print(header, breaks, recipe, breaks, recipe.sum("TOTAL"))
}

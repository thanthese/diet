package main

import "fmt"

func main() {
	recipe := ingredients{
		spinach(85),
		eggEach(1),
		eggEach(2),
	}.sort(byNetCarbs)
	fmt.Print(header, breaks, recipe, breaks, recipe.sum("TOTAL"))
}

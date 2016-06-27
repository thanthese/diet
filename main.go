package main

import "fmt"

func main() {
	recipe := ingredients{
		eggEach(3),
		spinach(100),
	}.sort(byNetCarbs)
	fmt.Print(header, breaks, recipe, breaks, recipe.sum("TOTAL"))
}

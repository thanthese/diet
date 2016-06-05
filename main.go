package main

import "fmt"

func main() {
	printHeader()
	printBreaks()
	smoothie().sortByProtein().reverse().print()
	printBreaks()
	smoothie().sum("SMOOTHIE").print()
}

type food struct {
	name     string
	dollars  float32
	g        float32
	carbs    float32
	protein  float32
	fat      float32
	fiber    float32
	calories float32
}

func printHeader() { fmt.Println("INGREDIENT        NET CARBS   PROTEIN   CALORIES       $      G") }
func printBreaks() { fmt.Println("---------------------------------------------------------------") }

func (f food) netCarbs() float32 {
	return f.carbs - f.fiber
}

func (f food) print() {
	fmt.Printf("%15s %11.0f %9.0f %10.0f %7.2f %6.0f\n",
		f.name,
		f.netCarbs(),
		f.protein,
		f.calories,
		f.dollars,
		f.g,
	)
}

type foods []food

func (fs foods) sum(name string) (total food) {
	total.name = name
	for _, f := range fs {
		total.dollars += f.dollars
		total.g += f.g
		total.carbs += f.carbs
		total.protein += f.protein
		total.fat += f.fat
		total.fiber += f.fiber
		total.calories += f.calories
	}
	return
}

func (fs foods) print() {
	for _, f := range fs {
		f.print()
	}
}

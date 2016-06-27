package main

import "fmt"

const (
	header = "INGREDIENT        NET CARBS   PROTEIN   CALORIES       $   FIBER    AMT\n"
	breaks = "-----------------------------------------------------------------------\n"
)

type ingredient struct {
	name     string
	dollars  float64
	carbs    float64
	protein  float64
	fat      float64
	fiber    float64
	calories float64
	factor   float64
}

func (ing ingredient) netCarbs() float64 {
	return ing.carbs - ing.fiber
}

func (in ingredient) times(factor float64) (out ingredient) {
	out.name = in.name
	out.dollars = in.dollars * factor
	out.carbs = in.carbs * factor
	out.protein = in.protein * factor
	out.fat = in.fat * factor
	out.fiber = in.fiber * factor
	out.calories = in.calories * factor
	out.factor = factor
	return
}

func (ing ingredient) String() string {
	return fmt.Sprintf("%15s %11.0f %9.0f %10.0f %7.2f %7.0f %6.0f\n",
		ing.name,
		ing.netCarbs(),
		ing.protein,
		ing.calories,
		ing.dollars,
		ing.fiber,
		ing.factor,
	)
}

type ingredients []ingredient

func (fs ingredients) sum(name string) (total ingredient) {
	total.name = name
	for _, f := range fs {
		total.dollars += f.dollars
		total.carbs += f.carbs
		total.protein += f.protein
		total.fat += f.fat
		total.fiber += f.fiber
		total.calories += f.calories
	}
	total.factor = 0 // there's no sensible way to add 3 eggs to 80g spinach
	return
}

func (fs ingredients) String() (out string) {
	for _, f := range fs {
		out += f.String()
	}
	return
}

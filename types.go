package main

import "fmt"

const (
	header = "INGREDIENT        NET CARBS   PROTEIN   CALORIES       $    AMT\n"
	breaks = "---------------------------------------------------------------\n"
)

// foodData includes only basic information about the food that I found online.
//
// The units get tricky. Usually they're in terms of per gram, but some foods
// are per unit like eggs or fillets. In those cases the units are per count.
// The name of the food will give you a clue how it's used: look for keywords
// like "can" or "each".
//
// serving normally means number of grams in a serving. If it's 1 we're working
// with a per count food. If it's 0 I have no data on a typical serving.
type foodData struct {
	name     string
	serving  float64
	dollars  float64
	carbs    float64
	protein  float64
	fat      float64
	fiber    float64
	calories float64
}

func (f foodData) netCarbs() float64 {
	return f.carbs - f.fiber
}

// g returns an ingredient with this many grams of this food.
func (fd foodData) g(grams float64) (ing ingredient) {
	ing.name = fd.name
	ing.serving = 0 // anything else could be meaningless
	ing.dollars = fd.dollars * grams
	ing.carbs = fd.carbs * grams
	ing.protein = fd.protein * grams
	ing.fat = fd.fat * grams
	ing.fiber = fd.fiber * grams
	ing.calories = fd.calories * grams
	ing.amt = grams
	return
}

// servings returns an ingredient with this many servings of this food. For
// example, 2 servings of spinach instead of having to figure and write 170g of
// spinach.
func (fd foodData) servings(num float64) (ing ingredient) {
	return fd.g(fd.serving * num)
}

// count is the number of whatever, like 2 eggs.
func (fd foodData) count(num float64) (ing ingredient) {
	return fd.g(fd.serving * num)
}

// ingredient is just like a foodData, but with an amount used -- like I put
// 60g of spinach in my salad. Amt plays by all the same odd rules that serving
// does in foodData.
type ingredient struct {
	foodData
	amt float64
}

func (ing ingredient) String() string {
	return fmt.Sprintf("%15s %11.0f %9.0f %10.0f %7.2f %6.0f\n",
		ing.name,
		ing.netCarbs(),
		ing.protein,
		ing.calories,
		ing.dollars,
		ing.amt,
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
	total.amt = 0 // there's no sensible way to add 3 eggs to 80g spinach
	return
}

func (fs ingredients) String() (out string) {
	for _, f := range fs {
		out += f.String()
	}
	return
}

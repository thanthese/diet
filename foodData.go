package main

// nfs = nutrition facts serving

const gramsInAPound = 453.592

func dollarsPerPound(dollars float64) float64 {
	return dollars / gramsInAPound
}

//////////////////////////////////////////////////////////////////////

func spinach(factor float64) (ing ingredient) {
	nfs := 100.0
	return ingredient{
		name:     "spinach",
		dollars:  dollarsPerPound(4),
		calories: 23 / nfs,
		fat:      0 / nfs,
		carbs:    4 / nfs,
		fiber:    2 / nfs,
		protein:  3 / nfs,
	}.times(factor)
}

func broccoli(factor float64) (ing ingredient) {
	nfs := 100.0
	return ingredient{
		name:     "broccoli",
		dollars:  0,
		calories: 35 / nfs,
		fat:      0 / nfs,
		carbs:    7 / nfs,
		fiber:    3 / nfs,
		protein:  2 / nfs,
	}.times(factor)
}

func brusselsSprouts(factor float64) (ing ingredient) {
	nfs := 100.0
	return ingredient{
		name:     "b. sprouts",
		dollars:  0,
		calories: 36 / nfs,
		fat:      0 / nfs,
		carbs:    7 / nfs,
		fiber:    3 / nfs,
		protein:  3 / nfs,
	}.times(factor)
}

func avocado(factor float64) (ing ingredient) {
	nfs := 100.0
	return ingredient{
		name:     "avocado",
		dollars:  1 / 201.0,
		calories: 160 / nfs,
		fat:      15 / nfs,
		carbs:    9 / nfs,
		fiber:    7 / nfs,
		protein:  2 / nfs,
	}.times(factor)
}

func hempSeeds(factor float64) (ing ingredient) {
	nfs := 30.0
	return ingredient{
		name:     "hemp seeds",
		dollars:  12 / (nfs * 26),
		calories: 170 / nfs,
		fat:      13 / nfs,
		carbs:    3 / nfs,
		fiber:    3 / nfs,
		protein:  10 / nfs,
	}.times(factor)
}

func onion(factor float64) (ing ingredient) {
	nfs := 100.0
	return ingredient{
		name:     "onion",
		dollars:  0,
		calories: 42 / nfs,
		fat:      0 / nfs,
		carbs:    10 / nfs,
		fiber:    1 / nfs,
		protein:  1 / nfs,
	}.times(factor)
}

func redPepper(factor float64) (ing ingredient) {
	nfs := 100.0
	return ingredient{
		name:     "red pepper",
		dollars:  0,
		calories: 26 / nfs,
		fat:      0 / nfs,
		carbs:    6 / nfs,
		fiber:    1 / nfs,
		protein:  1 / nfs,
	}.times(factor)
}

func chiaSeeds(factor float64) (ing ingredient) {
	nfs := 12.0
	return ingredient{
		name:     "chia seeds",
		dollars:  12 / (nfs * 75),
		calories: 60 / nfs,
		fat:      3 / nfs,
		carbs:    5 / nfs,
		fiber:    5 / nfs,
		protein:  3 / nfs,
	}.times(factor)
}

// very unscientific
func cheese(factor float64) (ing ingredient) {
	nfs := 100.0
	return ingredient{
		name:     "cheese",
		dollars:  0,
		calories: 375 / nfs,
		fat:      31 / nfs,
		carbs:    2 / nfs,
		fiber:    0 / nfs,
		protein:  22 / nfs,
	}.times(factor)
}

func almonds(factor float64) (ing ingredient) {
	nfs := 100.0
	return ingredient{
		name:     "almonds",
		dollars:  dollarsPerPound(6),
		calories: 575 / nfs,
		fat:      49 / nfs,
		carbs:    22 / nfs,
		fiber:    12 / nfs,
		protein:  21 / nfs,
	}.times(factor)
}

func oliveOil(factor float64) (ing ingredient) {
	nfs := 14.0
	return ingredient{
		name:     "olive oil",
		dollars:  12 / (nfs * 109),
		calories: 130 / nfs,
		fat:      14 / nfs,
		carbs:    0 / nfs,
		fiber:    0 / nfs,
		protein:  0 / nfs,
	}.times(factor)
}

func coconutOil(factor float64) (ing ingredient) {
	nfs := 100.0
	return ingredient{
		name:     "coconut oil",
		dollars:  0,
		calories: 862 / nfs,
		fat:      100 / nfs,
		carbs:    0 / nfs,
		fiber:    0 / nfs,
		protein:  0 / nfs,
	}.times(factor)
}

func almondButter(factor float64) (ing ingredient) {
	nfs := 32.0
	return ingredient{
		name:     "almond butter",
		dollars:  10 / (nfs * 24),
		calories: 210 / nfs,
		fat:      18 / nfs,
		carbs:    7 / nfs,
		fiber:    3 / nfs,
		protein:  6 / nfs,
	}.times(factor)
}

func bananas(factor float64) (ing ingredient) {
	nfs := 100.0
	return ingredient{
		name:     "bananas",
		dollars:  0,
		calories: 89 / nfs,
		fat:      0 / nfs,
		carbs:    23 / nfs,
		fiber:    3 / nfs,
		protein:  1 / nfs,
	}.times(factor)
}

func lentils(factor float64) (ing ingredient) {
	// http://nutritiondata.self.com/facts/legumes-and-legume-products/4338/2
	nfs := 100.0
	return ingredient{
		name:     "lentils",
		dollars:  (2.79 / gramsInAPound) / 3,
		calories: 116 / nfs,
		fat:      0 / nfs,
		carbs:    20 / nfs,
		fiber:    8 / nfs,
		protein:  9 / nfs,
	}.times(factor)
}

func apples(factor float64) (ing ingredient) {
	// http://nutritiondata.self.com/facts/fruits-and-fruit-juices/1809/2
	nfs := 100.0
	return ingredient{
		name:     "apples",
		dollars:  2.50 / gramsInAPound,
		calories: 52 / nfs,
		fat:      0 / nfs,
		carbs:    14 / nfs,
		fiber:    2 / nfs,
		protein:  0 / nfs,
	}.times(factor)
}

func pintoBeans(factor float64) (ing ingredient) {
	// http://nutritiondata.self.com/facts/legumes-and-legume-products/4430/2
	nfs := 100.0
	return ingredient{
		name:     "pinto beans",
		dollars:  2 / gramsInAPound,
		calories: 143 / nfs,
		fat:      1 / nfs,
		carbs:    26 / nfs,
		fiber:    9 / nfs,
		protein:  9 / nfs,
	}.times(factor)
}

func brownRice(factor float64) (ing ingredient) {
	// http://nutritiondata.self.com/facts/cereal-grains-and-pasta/5707/2
	nfs := 100.0
	return ingredient{
		name:     "brown rice",
		dollars:  2 / gramsInAPound,
		calories: 111 / nfs,
		fat:      1 / nfs,
		carbs:    23 / nfs,
		fiber:    2 / nfs,
		protein:  3 / nfs,
	}.times(factor)
}

func strawberries(factor float64) (ing ingredient) {
	// http://nutritiondata.self.com/facts/fruits-and-fruit-juices/2064/2
	nfs := 100.0
	return ingredient{
		name:     "strawberries",
		dollars:  2.50 / gramsInAPound,
		calories: 32 / nfs,
		fat:      0 / nfs,
		carbs:    8 / nfs,
		fiber:    2 / nfs,
		protein:  1 / nfs,
	}.times(factor)
}

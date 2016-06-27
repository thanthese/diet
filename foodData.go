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
		name:     "brussels sprouts",
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

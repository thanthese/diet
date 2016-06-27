package main

func eggEach(factor float64) (ing ingredient) {
	return ingredient{
		name:     "egg",
		dollars:  4 / 12.0,
		calories: 70,
		fat:      5,
		carbs:    0,
		fiber:    0,
		protein:  6,
	}.times(factor)
}

func sardinesCanBonesEach(factor float64) (ing ingredient) {
	nfs := 2.5
	return ingredient{
		name:     "sardines can",
		dollars:  24 / 12.0,
		calories: 110 * nfs,
		fat:      8 * nfs,
		carbs:    0 * nfs,
		fiber:    0 * nfs,
		protein:  10 * nfs,
	}.times(factor)
}

func salmonFilletEach(factor float64) (ing ingredient) {
	return ingredient{
		name:     "salmon",
		dollars:  20 / 7,
		calories: 340,
		fat:      20,
		carbs:    0,
		fiber:    0,
		protein:  37,
	}.times(factor)
}

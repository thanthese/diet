package main

// Everything is in terms of "per 100 g" (except for g, which is the amount of
// that thing that you're eating).

// perPound takes the price per pound and returns the price per 100 grams.
func perPound(per100g float32) float32 {
	return per100g * 0.220462
}

////////////////////////////////////////////////////////////////////////////////

func water(g float32) food {
	return food{
		name:     "water",
		g:        g,
		dollars:  0,
		carbs:    0,
		protein:  0,
		fat:      0,
		fiber:    0,
		calories: 0,
	}
}

func lentils(g float32) food {
	return food{
		name:     "lentils",
		g:        g,
		dollars:  0.01 * g * perPound(2.39) / 5,
		carbs:    0.01 * g * 20.13,
		protein:  0.01 * g * 9.02,
		fat:      0.01 * g * 0.38,
		fiber:    0.01 * g * 7.9,
		calories: 0.01 * g * 116,
	}
}

func spinach(g float32) food {
	return food{
		name:     "spinach",
		g:        g,
		dollars:  0.01 * g * perPound(4.29),
		carbs:    0.01 * g * 3.63,
		protein:  0.01 * g * 2.86,
		fat:      0.01 * g * 0.39,
		fiber:    0.01 * g * 2.2,
		calories: 0.01 * g * 23,
	}
}

func strawberries(g float32) food {
	return food{
		name:     "strawberries",
		g:        g,
		dollars:  0.01 * g * perPound(1.67),
		carbs:    0.01 * g * 9.13,
		protein:  0.01 * g * 0.43,
		fat:      0.01 * g * 0.11,
		fiber:    0.01 * g * 2.1,
		calories: 0.01 * g * 35,
	}
}

func almonds(g float32) food {
	return food{
		name:     "almonds",
		g:        g,
		dollars:  0.01 * g * perPound(5.63),
		carbs:    0.01 * g * 21.55,
		protein:  0.01 * g * 21.15,
		fat:      0.01 * g * 49.93,
		fiber:    0.01 * g * 12.5,
		calories: 0.01 * g * 579,
	}
}

func hempSeeds(g float32) food {
	return food{
		name:     "hemp seeds",
		g:        g,
		dollars:  0.01 * g * (13.69 / 793) * 100,
		carbs:    0.01 * g * 8.67,
		protein:  0.01 * g * 31.56,
		fat:      0.01 * g * 48.75,
		fiber:    0.01 * g * 4,
		calories: 0.01 * g * 553,
	}
}

func applesGala(g float32) food {
	return food{
		name:     "apples, gala",
		g:        g,
		dollars:  0.01 * g * perPound(2.50),
		carbs:    0.01 * g * 13.68,
		protein:  0.01 * g * 0.25,
		fat:      0.01 * g * 0.12,
		fiber:    0.01 * g * 2.3,
		calories: 0.01 * g * 57,
	}
}

package main

// All units are in terms of per gram. Unless the name contains "serving"
// "can", "each", or similar, then it's per serving.

// nfs = nutrition facts serving

const gramsInAPound = 453.592

func spinachRaw() (fd foodData) {
	nfs := 100.0
	fd.name = "spinach, raw"
	fd.serving = 85
	fd.dollars = 4 / gramsInAPound
	fd.calories = 23 / nfs
	fd.fat = 0 / nfs
	fd.carbs = 4 / nfs
	fd.fiber = 2 / nfs
	fd.protein = 3 / nfs
	return
}

func eggEach() (fd foodData) {
	fd.name = "egg"
	fd.serving = 1
	fd.dollars = 4 / 12.0
	fd.calories = 70
	fd.fat = 5
	fd.carbs = 0
	fd.fiber = 0
	fd.protein = 6
	return
}

func hempSeeds() (fd foodData) {
	fd.name = "hemp seeds"
	fd.serving = 30
	fd.dollars = 12 / (fd.serving * 26)
	fd.calories = 170 / fd.serving
	fd.fat = 13 / fd.serving
	fd.carbs = 3 / fd.serving
	fd.fiber = 3 / fd.serving
	fd.protein = 10 / fd.serving
	return
}

func chiaSeeds() (fd foodData) {
	fd.name = "chia seeds"
	fd.serving = 12
	fd.dollars = 12 / (fd.serving * 75)
	fd.calories = 60 / fd.serving
	fd.fat = 3 / fd.serving
	fd.carbs = 5 / fd.serving
	fd.fiber = 5 / fd.serving
	fd.protein = 3 / fd.serving
	return
}

func coconutOil() (fd foodData) {
	fd.name = "coconut oil"
	fd.serving = 14
	fd.dollars = 12 / (fd.serving * 109)
	fd.calories = 130 / fd.serving
	fd.fat = 14 / fd.serving
	fd.carbs = 0 / fd.serving
	fd.fiber = 0 / fd.serving
	fd.protein = 0 / fd.serving
	return
}

func oliveOil() (fd foodData) {
	fd.name = "olive oil"
	fd.serving = 15
	fd.dollars = 10 / (fd.serving * 67)
	fd.calories = 120 / fd.serving
	fd.fat = 14 / fd.serving
	fd.carbs = 0 / fd.serving
	fd.fiber = 0 / fd.serving
	fd.protein = 0 / fd.serving
	return
}

func almondButter() (fd foodData) {
	fd.name = "almond butter"
	fd.serving = 32
	fd.dollars = 10 / (fd.serving * 24)
	fd.calories = 210 / fd.serving
	fd.fat = 18 / fd.serving
	fd.carbs = 7 / fd.serving
	fd.fiber = 3 / fd.serving
	fd.protein = 6 / fd.serving
	return
}

func sardinesCan() (fd foodData) {
	nfs := 2.5
	fd.name = "sardines can"
	fd.serving = 1
	fd.dollars = 24 / 12.0
	fd.calories = 110 * nfs
	fd.fat = 8 * nfs
	fd.carbs = 0 * nfs
	fd.fiber = 0 * nfs
	fd.protein = 10 * nfs
	return
}

func cucumber() (fd foodData) {
	nfs := 100.0
	fd.name = "cucumber"
	fd.serving = 0
	fd.dollars = 3
	fd.calories = 15 / nfs
	fd.fat = 0 / nfs
	fd.carbs = 4 / nfs
	fd.fiber = 0 / nfs
	fd.protein = 1 / nfs
	return
}

func salmonFillet() (fd foodData) {
	fd.name = "salmon"
	fd.serving = 1
	fd.dollars = 20 / 7
	fd.calories = 340
	fd.fat = 20
	fd.carbs = 0
	fd.fiber = 0
	fd.protein = 37
	return
}

func miniPeppersEach() (fd foodData) {
	nfs := 3.0
	fd.name = "mini peppers"
	fd.serving = 1
	fd.dollars = (10 / 6.0) / nfs
	fd.calories = 25 / nfs
	fd.fat = 0 / nfs
	fd.carbs = 5 / nfs
	fd.fiber = 1 / nfs
	fd.protein = 1 / nfs
	return
}

func ranch() (fd foodData) {
	fd.name = "ranch"
	fd.serving = 30
	fd.dollars = 5 / (16.0 * fd.serving)
	fd.calories = 150 / fd.serving
	fd.fat = 16 / fd.serving
	fd.carbs = 2 / fd.serving
	fd.fiber = 0 / fd.serving
	fd.protein = 0 / fd.serving
	return
}

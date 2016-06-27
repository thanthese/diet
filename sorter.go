package main

import "sort"

type cmp func(ingredient, ingredient) bool

type sortable struct {
	foods ingredients
	cmp   cmp
}

func (s sortable) Len() int           { return len(s.foods) }
func (s sortable) Swap(i, j int)      { s.foods[i], s.foods[j] = s.foods[j], s.foods[i] }
func (s sortable) Less(i, j int) bool { return s.cmp(s.foods[i], s.foods[j]) }

func (ings ingredients) sort(fn cmp) ingredients {
	sort.Sort(sortable{ings, fn})
	return ings
}

func byName(a, b ingredient) bool     { return a.name < b.name }
func byDollars(a, b ingredient) bool  { return a.dollars > b.dollars }
func byG(a, b ingredient) bool        { return a.factor > b.factor }
func byNetCarbs(a, b ingredient) bool { return a.netCarbs() > b.netCarbs() }
func byProtein(a, b ingredient) bool  { return a.protein > b.protein }
func byFat(a, b ingredient) bool      { return a.fat > b.fat }
func byFiber(a, b ingredient) bool    { return a.fiber > b.fiber }
func byCalories(a, b ingredient) bool { return a.calories > b.calories }

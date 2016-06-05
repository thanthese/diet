package main

import "sort"

type sortable struct {
	foods foods
	less  func(int, int) bool
}

func (s sortable) Len() int           { return len(s.foods) }
func (s sortable) Swap(i, j int)      { s.foods[i], s.foods[j] = s.foods[j], s.foods[i] }
func (s sortable) Less(i, j int) bool { return s.less(i, j) }

func (fs foods) sortByName() foods {
	sort.Sort(sortable{fs, func(i, j int) bool { return fs[i].name < fs[j].name }})
	return fs
}

func (fs foods) sortByDollars() foods {
	sort.Sort(sortable{fs, func(i, j int) bool { return fs[i].dollars < fs[j].dollars }})
	return fs
}

func (fs foods) sortByG() foods {
	sort.Sort(sortable{fs, func(i, j int) bool { return fs[i].g < fs[j].g }})
	return fs
}

func (fs foods) sortByNetCarbs() foods {
	sort.Sort(sortable{fs, func(i, j int) bool { return fs[i].netCarbs() < fs[j].netCarbs() }})
	return fs
}

func (fs foods) sortByProtein() foods {
	sort.Sort(sortable{fs, func(i, j int) bool { return fs[i].protein < fs[j].protein }})
	return fs
}

func (fs foods) sortByFat() foods {
	sort.Sort(sortable{fs, func(i, j int) bool { return fs[i].fat < fs[j].fat }})
	return fs
}

func (fs foods) sortByFiber() foods {
	sort.Sort(sortable{fs, func(i, j int) bool { return fs[i].fiber < fs[j].fiber }})
	return fs
}

func (fs foods) sortByCalories() foods {
	sort.Sort(sortable{fs, func(i, j int) bool { return fs[i].calories < fs[j].calories }})
	return fs
}

func (fs foods) reverse() foods {
	// https://github.com/golang/go/wiki/SliceTricks
	for i := len(fs)/2 - 1; i >= 0; i-- {
		opp := len(fs) - 1 - i
		fs[i], fs[opp] = fs[opp], fs[i]
	}
	return fs
}

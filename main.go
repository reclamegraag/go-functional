package main

import (
	. "functional/functions"
	. "strings"
)

func main() {
	var a StringSlice = []string{"x", "z", "z", "y"}
	var r StringSlice = []string{"a", "a", "z", "y"}
	var x IntSlice = []int{1, 1, 2, 3, 4}
	//var q IntSlice = []int{1, 2, 2, 5, 6}
	b := a.Unique()
	y := x.Unique()
	println(b[0], b[1], b[2])
	println(y[0], y[1], y[2], y[3])
	println(contains1(1))

	s := a.Intersect(r)
	println("intersect: ", s[0], s[1])
	var filtered = x.Filter(contains1)
	println(len(filtered))
	var i = IntSlice{1, 2, 3}
	c := a.Map(func(b string) string {
		return b + "c"
	})
	d := a.Map(ToUpper)
	println(c[0] + d[0])
	j := i.Map(func(x int) int { return x + 1 })
	println(j[0])
}

func contains1(i int) bool {
	isOne := i == 5
	return isOne
}

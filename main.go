package main

import (
	. "functional/functions"
	"strconv"
	. "strings"
)

func main() {
	Timer(func() {
		var list = []string{}
		for i := 0; i < 100000000; i++ {
			list = append(list, strconv.Itoa(i))
		}
		println(list[10000])
	})

	var a StringSlice = []string{"x", "z", "z", "y"}
	var r StringSlice = []string{"a", "a", "z", "y"}
	var x IntSlice = []int{1, 1, 2, 3, 4}
	//var q IntSlice = []int{1, 2, 2, 5, 6}

	a.Foreach(func(value string) {
		println(value)
	})

	x.Foreach(func(i int) {
		println(i)
	})

	b := a.Unique()
	y := x.Unique()
	println(b[0], b[1], b[2])
	println(y[0], y[1], y[2], y[3])
	println(contains1(1))

	s := a.Intersect(r)
	println("intersect: ", s[0], s[1])
	var filtered = x.Filter(contains1)
	var dropped = x.DropWhile(contains1)
	dropped.Foreach(func(i int) {
		println(i)
	})
	println(len(filtered))
	println("dropped: ", dropped[0])
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
	isOne := i == 1
	return isOne
}

package main

import (
	. "fp/functions"
	. "strings"
)

func main() {
	var a StringSlice = []string{"x", "z","z", "y"}
	var x IntSlice = []int{1,1,2,3,4}
	b := a.Unique()
	y := x.Unique()
	println(b[0],b[1],b[2])
	println(y[0],y[1],y[2],y[3])
	var i = IntSlice{1, 2, 3}
	c := a.Map(func(b string) string {
		return b + "c"
	})
	d := a.Map(ToUpper)
	println(c[0] + d[0])
	j := i.Map(func(x int) int { return x + 1 })
	println(j[0])
}

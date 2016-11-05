package main

import "fmt"

func main() {
	x := 5
	var a int
	var ergeb bool

	a, ergeb = half(x)
	fmt.Println(a)
	fmt.Println(ergeb)
}

func half(x int) (div int, moo bool) {
	div = x / 2
	if x%2 == 1 {
		moo = false
	} else {
		moo = true
	}
	return
}

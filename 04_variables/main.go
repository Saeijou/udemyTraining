package main

import "fmt"

func main() {

	var w int
	var x string
	var y float64
	var z bool

	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)

	fmt.Printf("%v \n", w)
	fmt.Printf("%v \n", x)
	fmt.Printf("%v \n", y)
	fmt.Printf("%v \n", z)
	fmt.Printf("%T \n", w)
	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", y)
	fmt.Printf("%T \n", z)


}

package main

import "fmt"

func zero(x *int) {
	fmt.Println(*x)
	fmt.Println(&x)
	*x = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
	fmt.Println(&x)
}

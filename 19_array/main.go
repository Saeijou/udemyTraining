package main

import "fmt"

func main() {
	var x [58]string
	//	fmt.Println(x)
	//	fmt.Println(len(x))
	//	fmt.Println(x[42])
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	var y [256]int

	for j := 0; j < 256; j++ {
		y[j] = j
	}

	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(y[42])


	var z [256]byte

	for j := 0; j < 256; j++ {
		z[j] = byte(j)
	}

	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(z[42])



}

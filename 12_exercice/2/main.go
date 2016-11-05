package main

import "fmt"

func main() {

	var one int
	var two int

	fmt.Println("Number 1")
	fmt.Scan(&one)
	fmt.Println("Number 2")
	fmt.Scan(&two)

	fmt.Printf("Remainder: %v\n", one%two)

	for i:=0; i<=100; i++{
		if i%2 == 0 {fmt.Printf("%v\n",i)}
	}

}

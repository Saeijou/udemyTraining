package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	name := "Stephan"

	fmt.Println("Hello, my nam is " + name)


	fmt.Println("What is your namee?")

	var yourname string

	fmt.Scan(&yourname)

	fmt.Println("Hello " + yourname)

}

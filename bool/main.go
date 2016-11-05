package main

import "fmt"

func main() {

	if true || false {
		fmt.Println("blabla")
	}

	if true && false {
		fmt.Println("blabla")
	}
}

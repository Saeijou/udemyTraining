package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning",
		1: "asdasd",
		2: "yfdsasd",
		3: "asdasd",
	}


	fmt.Println(myGreeting)
	//delete(myGreeting,2)

	if val, exists := myGreeting[2]; exists{
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else{
		fmt.Println("That value doesnt exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)

	for key, val := range myGreeting{
		fmt.Println(key, " - ", val)
	}

}

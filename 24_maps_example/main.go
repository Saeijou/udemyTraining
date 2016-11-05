package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	myGreeting2 := make(map[string]string)
	myGreeting["Tim"]= "Good morning."
	myGreeting["Jenny"]= "Bonjour"

	myGreetings := map[string]string{
		"Tim": "Good morning!",
		"Jenny": "Bonjour!",
	}
	myGreetings["Harleen"] = "Howdy"

	fmt.Println(myGreeting)
	fmt.Println(myGreeting2)
	fmt.Println(myGreetings)
	fmt.Println(len(myGreetings))
	delete(myGreetings, "Tim")
	fmt.Println(myGreetings)
}

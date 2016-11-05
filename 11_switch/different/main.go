package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("Contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("MyLeod")
	var t = contact{"Good to see you", "Maybe"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}

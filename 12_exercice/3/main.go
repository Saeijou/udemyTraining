package main

import "fmt"

func main() {

	var sum = 0

	for i := 0; i <= 1000; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Printf("%v FIZZBUZZ\n", i)
			sum = sum + i
		} else if i%3 == 0 {
			fmt.Printf("%v FIZZ\n", i)
			sum = sum + i
		} else if i%5 == 0 {
			fmt.Printf("%v BUZZ\n", i)
			sum = sum + i
		} else {
			fmt.Printf("%v\n", i)
		}

	}
	fmt.Println(sum)

}

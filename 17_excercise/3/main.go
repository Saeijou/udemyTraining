package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")

	var total int
	for i := 1; i < len(nums); i++ { //for _, v := range numbers
		if nums[i] > total {
			total = nums[i]
		}
	}

	fmt.Println(total)
}

func main() {

	sum(1, 2, 3, 4, 5)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	sum(1, 2, 3, 4, 5, 7 ,4, 1)
	sum()
}

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sumup(1, 10, 15, -15))
	fmt.Println(sumup(1, numbers...))
}

func sumup(startingValues int, numbers ...int) int {
	sum := startingValues
	for _, val := range numbers {
		sum += val
	}

	return sum
}

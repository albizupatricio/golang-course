package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{5, 1, 2}
	doubled := transformsNumbers(&numbers, double)
	tripled := transformsNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformedFunction(&numbers)
	transformerFn2 := getTransformedFunction(&moreNumbers)

	transformersNumbers := transformsNumbers(&numbers, transformerFn1)
	moreTransformersNumbers := transformsNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformersNumbers)
	fmt.Println(moreTransformersNumbers)
}

func transformsNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformedFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

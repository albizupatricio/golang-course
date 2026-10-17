package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	updatedPrices := append(prices, 5.99)
	updatedPrices[0] = 9.99
	updatedPrices[1] = 11.99

	fmt.Println(updatedPrices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	productNames[2] = "A Carpet"
// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(prices[3])
// 	fmt.Println(prices[1:3])
// }

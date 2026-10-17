package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {

	//1

	hobbies := [3]string{"Reading a book", "Playing football", "Drinking Vodka"}
	fmt.Println(hobbies)

	//2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	//3
	mainHobbies := hobbies[0:2]
	mainHobbies2 := hobbies[:2]
	fmt.Println(mainHobbies, mainHobbies2)

	//4
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	//5
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	//6
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Println(courseGoals)

	//7
	products := []Product{
		{"first-product", "A First Product", 12.99},
		{"second-product", "A Second Product", 129.99},
	}

	fmt.Println(products)

	newProduct := Product{
		"third-product",
		"A Third Product",
		15.99,
	}

	products = append(products, newProduct)
	fmt.Println(products)

}

// Time to practice what you learned!

// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

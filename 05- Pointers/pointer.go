package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("Age: ", age)
	fmt.Println("Age Pointer: ", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println("Age: ", age)

}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}

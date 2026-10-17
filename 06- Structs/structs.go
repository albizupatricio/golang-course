package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	admin := user.NewAdmin("test@example.com", "testpassword")

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputDetails()
	appUser.ClearUserName()
	appUser.OutputDetails()
	admin.User.OutputDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "01/01/1900",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("All fields are required.")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func (user *User) OutputDetails() {
	fmt.Println("User Details:")
	fmt.Println("First Name: ", user.firstName)
	fmt.Println("Last Name: ", user.lastName)
	fmt.Println("birthdate: ", user.birthdate)
	fmt.Println("Created At: ", user.createdAt.Format(time.RFC1123))
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

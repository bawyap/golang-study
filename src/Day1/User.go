package main

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	name  string
	age   int
	email string
}

func NewUser(name string, age int, email string) (*User, error) {
	if name == isEmpty() {
		//return nil, errors.New("user - name invalid")
		return nil, &ValidationError{Field: name, errMsg: "user - name Invalid"}
	}

	if (age < 0) || (age > 150) {
		return nil, errors.New("user - age invalid")
	}

	if !(strings.Contains(email, "@")) {
		//return nil, errors.New("email - invalid")
		return nil, &ValidationError{Field: email, errMsg: "user - email Invalid"}
	}

	return &User{
		name:  name,
		age:   age,
		email: email,
	}, nil
}

func (user User) isAdult() bool {
	return user.age >= 18
}

func main() {
	name := "Kanmani"
	age := 26
	email := "kanmanigmail.com"

	user, errr := NewUser(name, age, email)
	if errr != nil {
		fmt.Println("Error:", errr)
		return
	}

	fmt.Println("User Created:", name)
	fmt.Println("Is adult:", user.isAdult())
}

package main

import "fmt"

type UserData struct {
	id           uint
	name         string
	email        string
	password     string
	phone_number uint
}

func (u *UserData) setUserData(name string, email string, password string, phone_number uint) {
	u.id = 1
	u.name = name
	u.email = email
	u.password = password
	u.phone_number = phone_number
}

func main() {
	user := UserData{}
	var name string
	var email string
	var password string
	var phone_number uint

	fmt.Println("set user name:")
	fmt.Scanln(&name)
	fmt.Println("set user email:")
	fmt.Scanln(&email)
	fmt.Println("set user password:")
	fmt.Scanln(&password)
	fmt.Println("set user phone_number:")
	fmt.Scanln(&phone_number)

	user.setUserData(name, email, password, phone_number)

	fmt.Println(user)
}

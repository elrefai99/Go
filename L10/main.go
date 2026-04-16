package main

import "fmt"

type ISet interface {
	setData() string
}

type SetData struct {
	name         string
	email        string
	password     string
	username     string
	phone_number string
}

type IGet interface {
	getData() string
}
type GetData struct {
	name     string
	email    string
	username string
}

func (u *SetData) SetData(name string, email string, password string, phone_number string) {
	u.name = name
	u.email = email
	u.password = password
	u.username = name
	u.phone_number = phone_number
}

func (u *GetData) getData(name string, email string, username string) {
	u.name = name
}

func main() {
	user := SetData{}
	user2 := GetData{}
	user.SetData("Mohammed", "dadad", "dadad", "dadad")
	user2.getData("Mohammed", "dadad", "dadad")
	fmt.Println(user)
	fmt.Println(user2)
}

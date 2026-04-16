package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Users struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	users := []Users{
		{
			Name:  "Mohammed Mostafa",
			Email: "m.m@m.com",
		},
		{
			Name:  "Disha",
			Email: "disha@gmail.com",
		},
	}
	data, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = os.WriteFile("yyy.json", data, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = os.ReadFile("yyy.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Done...!")
}

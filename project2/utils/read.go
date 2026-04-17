package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Users struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ReadFilesJson() {
	dataJson, err := os.ReadFile("json/data.json")

	if err != nil {
		log.Fatal(err)
		return
	}
	var users []Users
	err = json.Unmarshal(dataJson, &users)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(users)
}

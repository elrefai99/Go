package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Write(newUsers Users) {

	dataJson, err := os.ReadFile("json/data.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	var user []Users
	if len(dataJson) > 0 {
		err = json.Unmarshal(dataJson, &user)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	if len(user) > 0 {
		newUsers.Id = user[len(user)-1].Id + 1
	} else {
		newUsers.Id = 1
	}
	user = append(user, newUsers)

	updated, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("json/data.json", updated, 0644)
	fmt.Println("User added:", newUsers)

}

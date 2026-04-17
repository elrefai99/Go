package main

import (
	"flag"
	"fmt"
	"log"
	"project/utils"
)

func main() {
	Read := flag.Bool("readFile", false, "read file")
	Write := flag.Bool("writeFile", false, "Write file")

	flag.Parse()

	if *Read {
		utils.ReadFilesJson()
		return
	}

	if *Write {
		user := utils.Users{}
		var name string
		var email string

		fmt.Println("set user name:")
		fmt.Scanln(&name)
		fmt.Println("set user email:")
		fmt.Scanln(&email)

		user.Name = name
		user.Email = email
		utils.Write(user)
		return
	}

	log.Fatal("Most add flag")
}

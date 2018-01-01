package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type userobject struct {
	Users []User `json:"users"`
}

// User object
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate int64  `json:"birth_date"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
}

func (u *userobject) Append(path string) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	var currentJson userobject

	e := json.Unmarshal(file, &currentJson)

	if e != nil {
		fmt.Println("json error Unmarshall")
		fmt.Println(e)
		os.Exit(1)
	}

	u.add(currentJson.Users)

	fmt.Printf("Размер исходного файла %v\n", len((*u).Users))
	fmt.Printf("Размер распарсеного файла %v\n", len(currentJson.Users))

	return

}

func (u *userobject) add(users []User) {
	u.Users = append(u.Users, users...)
}

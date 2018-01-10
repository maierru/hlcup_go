package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type visitobject struct {
	Visits []Visit `json:"visits"`
}

type Visit struct {
	User       int64 `json:"user"`
	Location   int64 `json:"location"`
	Visited_at int64 `json:"visited_at"`
	ID         int64 `json:"id"`
	Mark       uint8 `json:"mark"`
}

func (v *visitobject) Append(path string) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	var currentJson visitobject

	e := json.Unmarshal(file, &currentJson)

	if e != nil {
		fmt.Println("json error Unmarshall")
		fmt.Println(e)
		os.Exit(1)
	}

	v.add(currentJson.Visits)

	// fmt.Printf("Размер исходного файла %v\n", len((*v).Visits))
	// fmt.Printf("Размер распарсеного файла %v\n", len(currentJson.Visits))

	return
}

func (v *visitobject) add(visits []Visit) {
	v.Visits = append(v.Visits, visits...)
}

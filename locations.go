package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type jsonobject struct {
	Locations []Location `json:"locations"`
}

// Location object
type Location struct {
	Distance int64  `json:"distance"`
	City     string `json:"city"`
	Place    string `json:"place"`
	ID       int64  `json:"id"`
	Country  string `json:"country"`
}

// Append new locations to jsonobject
func (j *jsonobject) Append(path string) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	var currentJson jsonobject

	e := json.Unmarshal(file, &currentJson)

	if e != nil {
		fmt.Println("json error Unmarshall")
		os.Exit(1)
	}

	j.add(currentJson.Locations)

	fmt.Printf("Размер исходного файла %v\n", len((*j).Locations))
	fmt.Printf("Размер распарсеного файла %v\n", len(currentJson.Locations))

	return
}

func (j *jsonobject) add(l []Location) {
	j.Locations = append(j.Locations, l...)
}

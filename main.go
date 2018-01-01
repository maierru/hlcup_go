package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Started ")

	dir := "/home/kkk/srv/hlcupdocs/data/TRAIN/data/"

	files, er := ioutil.ReadDir(dir)

	if er != nil {
		fmt.Println("error dir read")
		os.Exit(1)
	}

	var locationJSON jsonobject
	var userJSON userobject
	var visitJSON visitobject

	for _, f := range files {
		if !f.IsDir() {
			name := f.Name()
			filename := strings.Split(name, ".")
			types := strings.Split(filename[0], "_")
			typeName := types[0]

			switch typeName {
			case "locations":
				file := dir + name
				locationJSON.Append(file)
				size := len(locationJSON.Locations)
				fmt.Printf("Обработали %v, размер locations %v \n", name, size)
			case "users":
				file := dir + name
				userJSON.Append(file)
				size := len(userJSON.Users)
				fmt.Printf("Обработали %v, размер users %v \n", name, size)
			case "visits":
				file := dir + name
				visitJSON.Append(file)
				size := len(visitJSON.Visits)
				fmt.Printf("Обработали %v, размер visits %v \n", name, size)
			default:
				fmt.Printf("Неизвестный тип файла %v\n", name)

			}

		}
	}
}

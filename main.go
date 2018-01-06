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
			case "users":
				file := dir + name
				userJSON.Append(file)
			case "visits":
				file := dir + name
				visitJSON.Append(file)
			default:
				fmt.Printf("Неизвестный тип файла: %v\n", name)

			}
		}
	}

	locations := make(map[int64]Location)
	users := make(map[int64]User)
	visits := make(map[int64]Visit)

	for _, l := range locationJSON.Locations {
		locations[l.ID] = l
	}

	for _, u := range userJSON.Users {
		users[u.ID] = u
	}

	for _, v := range visitJSON.Visits {
		visits[v.ID] = v
	}
	fmt.Printf("размер locations %v \n", len(locations))
	fmt.Printf("размер users %v \n", len(users))
	fmt.Printf("размер visits %v \n", len(visits))
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	// read file
	data, err := ioutil.ReadFile("./hello.json")
	if err != nil {
		fmt.Print(err)
	}

	// define data structure
	type DayPrice struct {
		Monitor_id int
		Name       string
		GBP        int
	}

	// json data
	var obj []DayPrice

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	//fmt.Printf("Birds : %+v", obj)

	for i := range obj {
		fmt.Println(obj[i].Monitor_id)
	}

}

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

	var obj DayPrice

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	// can access using struct now
	fmt.Printf("USD : %v\n", obj.Monitor_id)
	fmt.Println("EUR : " + obj.Name)
	fmt.Printf("GBP : %v\n", obj.GBP)

}

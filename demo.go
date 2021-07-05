package main

import (
	"encoding/json"
	"fmt"
)

// declaring a struct
type Country struct {

	// defining struct variables
	monitor_id int `json:monitor_id`
	Monitor_id int `json:Monitor_id`
}

// main function
func main() {

	// defining a struct instance
	//var country []Country

	// JSON array to be decoded
	// to an array in golang
	Data := Country{Monitor_id: 876787, monitor_id: 79798798}

	// decoding JSON array to
	// the country array
	bute = json.Marshal(Data)

	//fmt.Println(country)
	// printing decoded array
	// values one by one

	fmt.Println(string(bute))
	//fmt.Println(country[0].Monitor_id)

	//	for i := range country {
	//		fmt.Println(country[i].Monitor_id)
	//fmt.Println(country[i].Monitor_id)
	//	}
}

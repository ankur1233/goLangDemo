package demo

import (
	"encoding/json"
	"log"
)

// declaring a struct
type Country struct {
	Nonitor_id string
	Monitor_id string
}

func main() {

	var country []Country
	Data := []byte(`[
	{"nonitor_id":"2323","Monitor_id":"0909"},
	{"nonitor_id":"323","Monitor_id":"8888"}
	]`)
	json.Unmarshal(Data, &country)
	for i := range country {
		//fmt.Println(country[i].Nonitor_id + country[i].Monitor_id)
		log.Println(country[i].Nonitor_id + "<-->" + country[i].Monitor_id)
	}
}

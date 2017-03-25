package main

import "fmt"

func main() {
	//map is a data Type
	//with map datatype value can be stored corresponding to key.
	//var map_variable map[key_data_type]value_data_type

	var country_Capital map[string]string
	//Assign value or create Map
	country_Capital = make(map[string]string)

	//Insert key values into map
	country_Capital["France"] = "Paris"
	country_Capital["Italy"] = "Rome"
	country_Capital["India"] = "New Delhi"
	//  Country_Capital["United States"] = "Washington DC"

	//Print Map using key
	for country := range country_Capital {
		fmt.Println("Capital of", country, "is", country_Capital[country])

	}

	//Test if entry is present in Map
	capital, ok := country_Capital["United States"]
	if ok {
		fmt.Println("Capital of United states is", capital)
	} else {
		fmt.Println("Entry not present in map")
	}

}

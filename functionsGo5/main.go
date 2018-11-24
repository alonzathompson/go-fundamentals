package main

import "fmt"

func main() {

	//Map of maps
	//first map declares map
	//second map is a value for first map
	employees := map[string]map[string]string{
		"BT": map[string]string{
			"firstName": "Blake",
			"lastName":  "Travis",
		},
		"PC": map[string]string{
			"firstName": "Parker",
			"lastName":  "Cooper",
		},
		"LT": map[string]string{
			"firstName": "Larry",
			"lastName":  "Travis",
		},
	}

	// ok = employees with the initial PC
	// it is stored in emp
	// essentially if ok((if employees with name pc) print name with initials
	if emp, ok := employees["PC"]; ok {
		fmt.Println(emp["firstName"], emp["lastName"])
	}

	//loops through initials(key is index) and values which are employees(emp is value)
	for initials, emp := range employees {
		//print each initial and employee first name and last name
		fmt.Println(initials, emp["firstName"], emp["lastName"])
	}
}

package main

import "fmt"

// START OMIT
func main2() { // HL12
	type Profile struct {
		Name   string
		Age    int
		Skills []string
	}

	students := []Profile{
		{
			Name:   "Adehina",
			Age:    10,
			Skills: []string{"programming", "sleeping"},
		},
	}

	age := students[1].Age
	fmt.Println(age) // HL
} // HL12

//END OMIT

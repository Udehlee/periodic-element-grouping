package main

import (
	"fmt"
)

// Element interface defines the  method for getting element details
type Element interface {
	GetElements() string
}

// ElementGroup represents a group of elements
type ElementGroup struct {
	Name     string
	Elements []Element
}

// ElementInfo struct holds information about an element
type ElementInfo struct {
	Name string
}

// GetDetails method for the reciever type ElementInfo to implement the Element interface
func (e ElementInfo) GetElements() string {
	return (e.Name)

}

func main() {
	// Create an array of ElementGroup
	groups := []ElementGroup{
		{"Group I", []Element{
			ElementInfo{"Hydrogen"},
			ElementInfo{"Lithium"},
			ElementInfo{"Sodium"},
			ElementInfo{"Potassium"},
		}},
		{"Group II", []Element{
			ElementInfo{"Beryllium"},
			ElementInfo{"Magnesium"},
			ElementInfo{"Calcium"},
		}},
		{"Group III", []Element{
			ElementInfo{"Boron"},
			ElementInfo{"Aluminium"},
		}},
		{"Group IV", []Element{
			ElementInfo{"Carbon"},
			ElementInfo{"silicon"},
		}},
		{"Group V", []Element{
			ElementInfo{"Nitrogen"},
			ElementInfo{"Phosphorus"},
		}},
		{"Group VI", []Element{
			ElementInfo{"Oxygen"},
			ElementInfo{"Sulphur"},
		}},
		{"Group VII", []Element{
			ElementInfo{"Flourine"},
			ElementInfo{"Chlorine"},
		}},
		{"Group O", []Element{
			ElementInfo{"Helium"},
			ElementInfo{"Neon"},
			ElementInfo{"Argon"},
		}},
	}

	// Print group name and  thier  elements
	//Range over the group ang access the the group name
	for _, group := range groups {
		fmt.Println(group.Name)
		for _, element := range group.Elements {
			//Range over each elements in each group and print the elements
			fmt.Println(element.GetElements())

		}
		fmt.Println("----------")
	}

}

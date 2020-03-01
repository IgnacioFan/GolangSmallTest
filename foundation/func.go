package main

import (
	"fmt"
)

func main () {
	fmt.Println(name("Nacho", "Fan"))
	// get multiple results
	fullname,err := nameWithErr("Ignatius", "Fan")
	// handle error
	if err != nil {
		fmt.Println("Handle Error case")
	}
	fmt.Println(fullname)
}

func name(firstName string, lastName string)(string) {
	fullname := firstName + " " + lastName
	return fullname
}
// multiple results
func nameWithErr(firstName string, lastName string)(string, error) {
	fullname := firstName + " " + lastName
	return fullname, nil
}
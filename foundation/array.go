package main

import (
	"fmt"
)

func main() {
	// declared an empty array of strings
	var days [2]string
	days[0] = "holiday"
	days[1] = "working day"
	fmt.Println(days)

	schedule := [4]string{"2/2", "2/8", "2/13", "2/25"}
	fmt.Println(schedule)
	
	// use range 
	for i, v := range days {
		fmt.Println(i, v)
	}

	// assigned an array with elements
	weekdays := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	fmt.Println(weekdays)
	// fmt.Println(weekdays[0])
	for i := 0; i < len(weekdays); i++ {
		fmt.Println(weekdays[i])
	}
}

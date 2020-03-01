package main

import (
	"fmt"
)

func main() {
	// slice is a dynamically-sized
	primes := [6]int{2,3,5,7,11,13}
	// include first ele, excludes last ele
	var s []int = primes[1:4] 
	fmt.Println(s)

	names := [4]string{
		"John", "Paul", "George", "Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a,b)

	// slice are like references to arrays
	b[0] = "xxx"
	fmt.Println(a,b)
	fmt.Println(names)

	// slice literal like array literal without length
	q := []int{2,3,5,7,11}
	fmt.Println(q)

	ss := []int{2,3,5,7,11,13}
	ss = ss[1:4]
	fmt.Println(ss)

	ss = ss[:2]
	fmt.Println(ss)

	ss = ss[1:]
	fmt.Println(ss)
}
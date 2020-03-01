package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

var (
	v1 = Vertex{3,4}
	v2 = Vertex{x:5}
	v3 = Vertex{}
	p = &Vertex{3,4} // has type *Vertex
)

func main() {
	// simple struct
	fmt.Println(Vertex{1,2})

	// access struct
	v := Vertex{1,2}
	v.x = 4
	fmt.Println(v)

	// struct with pointer
	a := &v
	a.x = 9
	fmt.Println(v)

	// struct literals
	fmt.Println(v1, p, v2, v3)
	
}
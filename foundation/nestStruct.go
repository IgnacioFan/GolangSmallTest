package main

import (
	"fmt"
)

func main() {
	type Person struct {
		name string
		age int
	}

	type Team struct {
		name string
		players [2]Person
	}

	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{
		Person{name: "John", age: 28},
		Person{name: "Paul", age: 30}}

	bigBang := Team{
		name: "BigBang",
		players: players}

	fmt.Println(bigBang)
}
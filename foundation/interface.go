package main

import(
	"fmt"
)

type Guitarist interface {
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar(){
	fmt.Printf("%s plays the Base \n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar(){
	fmt.Printf("%s plays the Acoustic \n", b.Name)
}

func main() {
	var player1 BaseGuitarist
	player1.Name = "Jeff"
	player1.PlayGuitar()
	
	var player2 AcousticGuitarist
	player2.Name = "Jann"
	player2.PlayGuitar()

	var guitarists []Guitarist
	guitarists = append(guitarists, player1)
	guitarists = append(guitarists, player2)
	fmt.Println(guitarists)
}
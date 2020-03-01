package main

import (
	"fmt"
	"encoding/json"
)

type Book struct {
	Title string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Developer bool `json:"is_developer"`
}

func main(){
	fmt.Println("Hello world")
	author := Author{Name: "Nacho", Age: 27, Developer: true}
	book := Book{Title: "Learning Go", Author: author}

	fmt.Printf("%+v\n", book)

	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}
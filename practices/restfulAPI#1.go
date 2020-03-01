package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "home is starting!")
}

func allArticles(w http.ResponseWriter, req *http.Request){
	articles := Articles{
		Article{Title:"Title1", Desc: "first desc", Content: "first content"},
		Article{Title:"Title2", Desc: "second desc", Content: "second content"},
	}
	fmt.Println("call all articles api")
	json.NewEncoder(w).Encode(articles)
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	handleRequest()
}
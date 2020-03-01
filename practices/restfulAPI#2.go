package main

import(
	"encoding/json"
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

type event struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
}

type allEvents []event

var events = allEvents{
	{
		Id: "1", Title: "test1", Desc: "this is the first test",
	},
	{
		Id: "2", Title: "test2", Desc: "this is the second test",
	},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome home!")
}

func getAllEvents(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(events)
	fmt.Println("call all events api")
}

func createEvent(w http.ResponseWriter, r *http.Request){
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "please enter a title and description")
	}
	// fit into the event struct
	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	// 201 created status code
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]
	
	for _, singleEvent := range events {
		if singleEvent.Id == eventId {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
	
}

func updateEvent(w http.ResponseWriter, r *http.Request){
	eventId := mux.Vars(r)["id"]
	var updateEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "please enter data with the title and description")
	}

	json.Unmarshal(reqBody, &updateEvent)

	for i, singleEvent := range events {
		if singleEvent.Id == eventId {
			singleEvent.Title = updateEvent.Title
			singleEvent.Desc = updateEvent.Desc
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request){
	eventId := mux.Vars(r)["id"]

	for i, singleEvent := range events {
		if singleEvent.Id == eventId {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event %v has been deleted", eventId)
		}
	}
}

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
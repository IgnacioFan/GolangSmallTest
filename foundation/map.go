package main

import ("fmt")

func main() {
	youtubeSubscribers := map[string]int {
		"sport": 234,
		"news": 3465,
	}
	fmt.Println(youtubeSubscribers["sport"])
}
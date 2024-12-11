package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {

	fmt.Println("Server has started...")
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/events", events)
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	s := "This is a live stream you are going to see many things"
	tokens := strings.Split(s, " ")

	for _, token := range tokens {
		content := fmt.Sprintf("date: %s\n\n", token)
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		// time delay
		time.Sleep(time.Millisecond * 420)
	}
	// fmt.Fprintln(w, tokens)
	
}

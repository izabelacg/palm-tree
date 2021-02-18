package main

import (
"fmt"
"log"
"net/http"
"os"
"strconv"
"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld: received a request")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	params, ok := r.URL.Query()["sleep"]

	if ok {
		secondsToSleep, err := strconv.Atoi(params[0])
		if err == nil {
			log.Printf( "Sleeping for: %d seconds", secondsToSleep)
			timeDuration := time.Duration(secondsToSleep)
			time.Sleep(time.Second * timeDuration)
		} else {
			log.Printf("Error getting number of seconds to sleep")
		}
	}
	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", target)
}

func main() {
	log.Print("helloworld: starting server...")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
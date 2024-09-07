package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
)

var randomNumber int
var step int = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/range", func(w http.ResponseWriter, r *http.Request) {
		Random(w, r, false)
	})

	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		value := r.FormValue("value")
		guess, _ := strconv.Atoi(value)
		var message string
		right := false
		if randomNumber == guess {
			message = "You guessed it right!"
			right = true
		} else {
			message = "Try again!"
		}
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, message)
		if right {
			Random(w, r, true)
		} else {
			fmt.Fprintf(w, message)
		}
	})
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func Random(w http.ResponseWriter, r *http.Request, right bool) {
	maxRange := 10 + step
	minRange := 0 + step/2
	step += rand.IntN(4)

	randomNumber = rand.IntN(maxRange - minRange)
	fmt.Println("Random number is ", randomNumber)

	w.Header().Set("Content-Type", "text/plain")
	if right {
		// fmt.Fprintf(w, "\n<script>htmx.trigger(htmx.find('#range'), 'load');</script>")
	} else {
		fmt.Fprintf(w, "Guess the number it falls between %d and %d ", minRange, maxRange)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// FizzbuzzServer does something
func FizzbuzzServer(w http.ResponseWriter, r *http.Request) {

	num, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/fizzbuzz/"))
	if err != nil {
		fmt.Fprint(w, "I don't understand that number")
		log.Printf("Recieved bad request %q", w)
		return
	}

	out := ""
	if num%3 == 0 {
		out += "Fizz"
	}
	if num%5 == 0 {
		out += "Buzz"
	}
	if out == "" {
		out = strconv.Itoa(num)
	}

	fmt.Fprintln(w, out)
	log.Printf("Recieved %q, returned %q", num, out)
}

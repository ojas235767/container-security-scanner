// main.go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Current Date and Time: %s", t)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

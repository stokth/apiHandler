package lesEvgen

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter int

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Counter = ", strconv.Itoa(counter))
	} else {
		fmt.Fprintln(w, "Method not allowed")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		counter++
		fmt.Fprintln(w, "Counter incremented")
	} else {
		fmt.Fprintln(w, "Method not allowed")
	}
}

func main1() {
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/post", PostHandler)
	http.ListenAndServe("localhost:8080", nil)
}

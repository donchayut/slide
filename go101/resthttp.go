package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	handler := grade

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func grade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	s := fmt.Sprintf("%s score is grade %s\n", r.RequestURI[1:], "A")
	io.WriteString(w, s)
}

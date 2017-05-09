package main

import (
	"net/http"
)

// type myHandler struct {
// 	greeting string
// }

// func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(fmt.Sprintf("%s World!", mh.greeting)))
// }

func main() {
	// http.Handle("/", &myHandler{greeting: "Hello"})
	// http.ListenAndServe(":8000", nil)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Shrishty!"))
	})
	http.ListenAndServe(":8000", nil)
}
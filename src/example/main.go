package main

import (
	"html/template"
	"fmt"
	"os"
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

	// http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Shrishty!"))
	// })
	// http.ListenAndServe(":8000", nil)

	templatestring := `Lemonade Stand Supply`
	t, err := template.New("title").Parse(templatestring)

	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, nil)

	if err != nil {
		fmt.Println(err)
	}
}
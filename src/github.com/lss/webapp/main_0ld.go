package main

import (
	"html/template"
	"net/http"
	"log"
)

func old_main()  {
	// Method 1
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
	// 	http.ServeFile(w, r,  "public" + r.URL.Path)

	// Method 2
	// f, err := os.Open("public" + r.URL.Path)

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Println(err)
	// }

	// defer f.Close()

	// var contentType string

	// switch {
	// case strings.HasSuffix(r.URL.Path, "css"):
	// 	contentType = "text/css"
	// case strings.HasSuffix(r.URL.Path, "html"):
	// 	contentType = "text/html"
	// case strings.HasSuffix(r.URL.Path, "png"):
	// 	contentType = "image/png"
	// default:
	// 	contentType = "text/plain"
	// }

	// w.Header().Add("Content-Type", contentType)

	// // copy the contents of the file to response writer
	// io.Copy(w, f)
	// })

	// Method 3
	// http.ListenAAndServe(":8000", http.FileServer(http.Dir("public")))

	// Method 4 with templates
	templates := old_populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]            // dont want the slash at the start
		t := templates.Lookup(requestedFile + ".html")

		if t != nil {
			err := t.Execute(w, nil)

			if err != nil {
				log.Println(err)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	})

	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
}

func old_populateTemplates() *template.Template  {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}


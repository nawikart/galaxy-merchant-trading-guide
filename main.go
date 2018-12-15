package main

import (
		"fmt"
		"html/template"
		"net/http"
		"./api"
		)


// Main function
func main() {

    // handle static files in "assets" folder
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
    
    // url Api for convertion between number and roman numerals
	http.HandleFunc("/convert", api.Api)	

    // frontend for user interface
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("templates/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, nil)
    })

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
    // http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
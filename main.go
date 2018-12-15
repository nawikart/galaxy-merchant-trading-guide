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
    
    // route Api for convertion between number and roman numerals
    http.HandleFunc("/conversion", api.Conversion)
    
    // route Api to showing the output for each queston
	http.HandleFunc("/answer", api.Intergalactic)

    // frontend for user interface
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var t, err = template.ParseFiles("templates/index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, nil)
    })

    fmt.Println("starting web server at http://localhost:8082/")
    http.ListenAndServe(":8082", nil)
}
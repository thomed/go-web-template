package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/* Server will be available on localhost:port */
var port = 4321

/*
 * Initialize request handlers and begin listening for requests.
 */
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/greetings", greetings)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))

	fmt.Printf("Server listening on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}

/*
 * Handles requests to the site root.
 */
func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	} else {
		_ = tmpl.Execute(w, nil)
	}
}

/*
 * Endpoint to display a greeting after form submission.
 * Redirect back to site root if request isn't a POST.
 */
func greetings(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		fmt.Printf("Name %s\n", name)
		tmpl, err := template.ParseFiles("greeting.html")
		if err != nil {
			fmt.Println(err)
		} else {
			_ = tmpl.Execute(w, name)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

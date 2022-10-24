package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title, Content string
}

// # Caching parsed templates
// store the ready-to-use template in a variable, and repeatedly use
// the same template each time you need to generate the output
var simpleTemplate *template.Template

func init() {
	simpleTemplate = template.Must(template.ParseFiles("simple.html"))
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	simplePage := Page{
		Title:   "Portfolio",
		Content: "Hi! I am Raihan. I'm a software Engineer.",
	}
	err := simpleTemplate.Execute(w, simplePage)
	if err != nil {
		log.Println(err, "Failed to execute html file")
	}
}

func main() {
	fmt.Println("starting server....")
	http.HandleFunc("/", displayPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err, "Failed to listen and serve to TCP")
	}
}

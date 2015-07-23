package main

import (
	"net/http"
	"html/template"
	"fmt"
	"time"
)	

var msgs []string

type Page struct { 
    Body  string
}

func timegetter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Here's your time:\n", time.Now().UTC())
}

/*func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}*/

func poster (w http.ResponseWriter, r *http.Request) {
	var p Page
	t, _ := template.ParseFiles("post.html")
	t.Execute(w, p)
	msgs = append(msgs, p.Body)
	fmt.Fprintln(w, msgs)
}

func allmessages (w http.ResponseWriter, r *http.Request) {
	for _,msg := range msgs {
		fmt.Fprintln(w, msg)
	}
}

func main(){
	http.HandleFunc("/time", timegetter)
	http.HandleFunc("/message", poster)
	http.HandleFunc("/get", allmessages)
	http.ListenAndServe(":8080", nil)
}

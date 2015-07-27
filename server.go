package main

import (
	"net/http"
	"html/template"
	"fmt"
	"time"
)	

var msgs []string

func timegetter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Here's your time:\n", time.Now().UTC())
}

func poster (w http.ResponseWriter, r *http.Request) {
	t, err1 := template.ParseFiles("post.html")
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
		for _ , msg := range msgs {
			fmt.Fprintf(w, "<ul>"+"<li>%s</li>"+"</ul>", msg)
		}
	err := t.Execute(w, msgs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if r.FormValue("body") != "" {
		msgs = append(msgs, r.FormValue("body"))
	}
}

func main(){
	http.HandleFunc("/time", timegetter)
	http.HandleFunc("/message", poster)
	http.ListenAndServe(":8080", nil)
}

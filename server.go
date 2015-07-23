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
	t, _ := template.ParseFiles("post.html")
	t.Execute(w, msgs)
	msgs = append(msgs, r.FormValue("body"))
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

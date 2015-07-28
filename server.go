package main

import (
	"net/http"
	"html/template"
	"fmt"
	"time"
	"sync"
)	

var  (msgs []string
	 rw *sync.RWMutex)

func timegetter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Here's your time:\n", time.Now().UTC())
}

func poster (w http.ResponseWriter, r *http.Request) {
	t, err1 := template.ParseFiles("post.html")
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	if r.FormValue("body") != ""{
		rw.Lock()
		msgs = append(msgs, r.FormValue("body"))
		rw.Unlock()
	}
	err := t.Execute(w, msgs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main(){
	http.HandleFunc("/time", timegetter)
	http.HandleFunc("/message", poster)
	http.ListenAndServe(":8080", nil)
}

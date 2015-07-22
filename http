package main

import (
	"net/http"
	"fmt"
	"time"
)	

var msgs = make([]string, 10)

func timegetter(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	t = t.UTC()
	fmt.Fprintln(w, "Here's your time:\n", t)
}

func poster (w http.ResponseWriter, r *http.Request) {
	msgs = append(msgs, r.URL.Path[6:])
}

func allmessages (w http.ResponseWriter, r *http.Request) {
	for _,msg := range msgs {
		fmt.Fprintln(w, msg)
	}
}

func main(){
	http.HandleFunc("/time", timegetter)
	http.HandleFunc("/post/", poster)
	http.HandleFunc("/get", allmessages)
	http.ListenAndServe(":8080", nil)
}

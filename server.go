package main

import (
	"net/http"
	"html/template"
	"fmt"
	"time"
	"sync"
	"log"
)	

type Message struct{
	msgs []string
	mu   sync.Mutex
	tmpl *template.Template
}  

func (m *Message) NewMessage (name string) *Message{ //инициализатор
	m.msgs = make([]string, 0, 100)
	var err error
	if m.tmpl, err = template.ParseFiles(name); err != nil {
		log.Fatal(err.Error())
	}
	return m
}

func (m *Message) Add (msg string){
	if  msg != ""{
		m.mu.Lock()
		m.msgs = append(m.msgs, msg)
		m.mu.Unlock()
	}
}
func (m *Message) Clear () {
	m.msgs = make([]string, 0, 100)
	return 
}

func (m *Message) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := m.tmpl.Execute(w, m.msgs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if r.Method == "POST" {
		if r.FormValue("name") != "" {
		m.Clear()
		}
		msg := r.FormValue("body")
		if msg != "" {
			m.Add(msg)
		}
}}

func main(){
	var m *Message 
	m.NewMessage("post.html") 
	log.Println(m)
	http.HandleFunc("/time", func (w http.ResponseWriter, r *http.Request) {fmt.Fprintln(w, time.Now().UTC() .Format("2006-01-02T15:04:05Z07:00"))})
	http.HandleFunc("/message", m.ServeHTTP)
	if err :=  http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err.Error())
	}
}

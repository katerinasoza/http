package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

type Message struct {
	msgs []string
	mu   sync.Mutex
	tmpl *template.Template
}

func NewMessage(name string) *Message {
	tmpl, err := template.ParseFiles(name)
	if err != nil {
		log.Fatal(err)
	}
	return &Message{
		msgs: make([]string, 0, 100),
		tmpl: tmpl,
	}
}

func (m *Message) Add(msg string) {
	if msg != "" {
		m.mu.Lock()
		m.msgs = append(m.msgs, msg)
		m.mu.Unlock()
	}
}
func (m *Message) Clear() {
	m.msgs = make([]string, 0, 100)
}

func (m *Message) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("check") != "" {
			m.Clear()
		}
		if msg := r.FormValue("body"); msg != "" {
			m.Add(msg)
		}
	}
	if err := m.tmpl.Execute(w, m.msgs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	m := NewMessage("post.html")
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, time.Now().UTC().Format("2006-01-02T15:04:05Z07:00"))
	})
	http.Handle("/message", m)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}

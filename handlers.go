package main

import (
	"net/http"
	"fmt"
	"html"
	"github.com/gorilla/mux"
	"encoding/json"
	"time"
)

func WH_Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func WH_TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func WH_TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func WH_CPUUsage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Stat()); err != nil {
		panic(err)
	}
}

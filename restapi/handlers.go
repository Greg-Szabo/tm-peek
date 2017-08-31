package restapi

import (
	"net/http"
	"fmt"
	"html"
	"encoding/json"
	"github.com/greg-szabo/tm-peek/cpu"
	"github.com/greg-szabo/tm-peek/io"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func cpuUsage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cpu.Stat()); err != nil {
		panic(err)
	}
}

func ioUsage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(io.Stat()); err != nil {
		panic(err)
	}
}

func tmUsage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("x"); err != nil {
		panic(err)
	}
}

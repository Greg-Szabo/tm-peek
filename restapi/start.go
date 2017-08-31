package restapi

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

var (
	tendermintAddress string
	router *mux.Router
)

func Start(address,tendermint string, separateRoutine bool) {
	tendermintAddress = tendermint
	if separateRoutine {
		go func() {
			log.Fatal(http.ListenAndServe(address, Router()))
		}()
	} else {
		log.Fatal(http.ListenAndServe(address, Router()))
	}
}

func init() {
	createRouter()
}

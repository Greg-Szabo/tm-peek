package restapi

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router()(*mux.Router) {
	return router
}

func createRouter() {

	router = mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return
}


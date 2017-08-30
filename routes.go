package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		WH_Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		WH_TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		WH_TodoShow,
	},
	Route{
		"CPUUsage",
		"GET",
		"/cpu",
		WH_CPUUsage,
	},
}

package restapi

import (
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []route{
	route{
		"Index",
		"GET",
		"/",
		index,
	},
	route{
		"CPUUsage",
		"GET",
		"/cpu",
		cpuUsage,
	},
	route{
		"IOUsage",
		"GET",
		"/io",
		ioUsage,
	},
}

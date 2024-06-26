package router

import (
	"net/http"
	"urlShortener/src/handlers"
)

//Route is a basic sturct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"404 Not Found",
		"POST",
		"/",
		ThrowBlankResponse,
	},
	Route{
		"shorten",
		"POST",
		"/shorten",
		handlers.ShortenUrl,
	},
	Route{
		"redirecturl",
		"GET",
		"/redirecturl/",
		handlers.RedirectUrl,
	},
	Route{
		"getmetrics",
		"GET",
		"/getmetrics",
		handlers.GetMetrics,
	},
}

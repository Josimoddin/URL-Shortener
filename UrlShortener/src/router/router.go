package router

import (
	"log"
	"net/http"
	"urlShortener/src/entities"
)

func NewRouter() {
	for _, route := range routes {
		if route.Method == "POST" {
			http.Handle("/api"+route.Pattern, PostMiddleware(http.HandlerFunc(route.HandlerFunc)))
		} else {
			http.Handle("/api"+route.Pattern, GetMiddleware(http.HandlerFunc(route.HandlerFunc)))
		}
	}
}

func PostMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type, Authorization,Auth")
		w.Header().Set("Access-Control-Expose-Headers", "Autho")
		w.Header().Set("Cache-Control", "no-cache,no-store")
		if req.Method != "POST" {
			log.Println(req.Method + "is called in " + req.URL.Path)
			entities.ThrowJSONResponse(entities.NotPostMethodResponse(), w)
			return
		}
		next.ServeHTTP(w, req)
	})
}

func GetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,content-type, Authorization,Auth")
		w.Header().Set("Access-Control-Expose-Headers", "Autho")
		w.Header().Set("Cache-Control", "no-cache,no-store")
		if req.Method != "GET" {
			log.Println(req.Method + "is called in " + req.URL.Path)
			entities.ThrowJSONResponse(entities.NotPostMethodResponse(), w)
			return
		}
		next.ServeHTTP(w, req)
	})
}

func ThrowBlankResponse(w http.ResponseWriter, req *http.Request) {
	entities.ThrowJSONResponse(entities.BlankPathCheckResponse(), w)
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)

// StartAPI server simple metric API
func StartAPI() {
	log.Printf("Server started")

	router := newRouter()

	log.Fatal(http.ListenAndServe(":3232", router))
}

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = []route{
	route{
		"index",
		"GET",
		"/",
		index,
	},

	route{
		"getHealth",
		strings.ToUpper("Get"),
		"/api/v1/health",
		getHealth,
	},

	route{
		"getMetric",
		strings.ToUpper("Get"),
		"/api/v1/metrics/{metric}",
		getMetric,
	},

	route{
		"getMetrics",
		strings.ToUpper("Get"),
		"/api/v1/metrics",
		getMetrics,
	},
}

type metric struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

var metrics = [2]metric{
	{"luck", 4.2},
	{"happiness", 15.0},
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func getMetric(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	var mx metric
	for _, m := range metrics {
		if m.Name == vars["metric"] {
			mx = m
			js, _ := json.Marshal(mx)
			w.Write(js)
		}
	}
}

func getMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var ms []map[string]string
	for _, m := range metrics {
		ms = append(ms, map[string]string{"name": m.Name})
	}
	js, _ := json.Marshal(ms)
	w.Write(js)
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

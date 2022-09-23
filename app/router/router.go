package router

import (
	"hats-for-parties/config"
	h "hats-for-parties/handlers"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type RouteHandlers interface {
	HealthCheckHandler(w http.ResponseWriter, r *http.Request)
	RouteNotFoundHandler(w http.ResponseWriter, r *http.Request)
	StartParty(w http.ResponseWriter, r *http.Request)
	EndParty(w http.ResponseWriter, r *http.Request)
}

var routeHandlers RouteHandlers = h.RouteHandlers

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "HealthCheck",
		Method:      http.MethodGet,
		Pattern:     "/ping",
		HandlerFunc: routeHandlers.HealthCheckHandler,
	},
	Route{
		Name:        "StartParty",
		Method:      http.MethodPost,
		Pattern:     "/start-party/{hatsNumber:[0-9]+}",
		HandlerFunc: routeHandlers.StartParty,
	},
	Route{
		Name:        "EndParty",
		Method:      http.MethodPost,
		Pattern:     "/end-party/{partyId}",
		HandlerFunc: routeHandlers.EndParty,
	},
}

// StartRequestHandler worker
func StartRequestHandler() {
	// Create the server
	r := buildRouter()

	// Power the server on
	log.Printf("API listening on port %d", config.ServiceConfig.Port)
	log.Fatal("FATAL Error: " + http.ListenAndServe(":"+strconv.Itoa(config.ServiceConfig.Port), r).Error())
}

func buildRouter() *mux.Router {
	r := mux.NewRouter()

	for _, route := range routes {
		r.Handle(route.Pattern, handlers.LoggingHandler(os.Stdout, route.HandlerFunc)).Methods(route.Method)
		log.Printf("Listening Route: %s %s", route.Method, route.Pattern)
	}
	r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(routeHandlers.RouteNotFoundHandler)))

	return r
}

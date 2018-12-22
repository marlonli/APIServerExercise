package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marlonli/APIServerExercise/pkg/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Create new routers
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
	"PersistMetadata",
	"POST",
	"/persist",
	handlers.PersistMetadata,
	},
	Route{
	"GetMetadata",
	"GET",
	"/metadata",
	handlers.GetMetadata,
	},
	Route{
	"SearchMetadata",
	"POST",
	"/search",
	handlers.SearchMetadata,
	},
}
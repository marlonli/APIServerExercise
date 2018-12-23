package routers

import (
	"github.com/marlonli/APIServerExercise/pkg/logger"
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
		"/v1/persist",
		logger.Logger(handlers.PersistMetadata, "Persist MetaData"),
	},
	Route{
		"GetMetadata",
		"GET",
		"/v1/metadata",
		logger.Logger(handlers.GetMetadata, "Get MetaData"),
	},
	Route{
		"SearchMetadata",
		"POST",
		"/v1/search",
		logger.Logger(handlers.SearchMetadata, "Search MetaData"),
	},
}

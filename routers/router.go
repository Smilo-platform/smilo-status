package routers

import (
	"net/http"
	"didux-status/handlers"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes []Route

// init method which will be automatically run 
func init() {

	/* init /status URL and handler function. */
	routes = append(routes, Route{
		Name:        "Status",
		Methods:     []string{"GET"},
		Pattern:     "/",
		HandlerFunc: handlers.HandleStatus,
	})
}

/* registering all url paths */
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Methods...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

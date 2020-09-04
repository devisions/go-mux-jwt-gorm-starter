package rest

import (
	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/routes"
	"github.com/devisions/go-mux-jwt-gorm-starter/users"
	"net/http"

	"github.com/gorilla/mux"
)

func NewApiRestRouter() *mux.Router {

	router := mux.NewRouter()

	// setup the routes
	var ApiRestRoutes []routes.ApiRestRouteSet
	ApiRestRoutes = append(ApiRestRoutes, users.UsersApiRestRoutes)

	for _, route := range ApiRestRoutes {
		subRouter := router.PathPrefix(route.Prefix).Subrouter()

		for _, sr := range route.SubRoutes {
			var handler http.Handler
			handler = sr.HandlerFunc
			if sr.Protected {
				handler = JWTMiddleware(sr.HandlerFunc)
			}
			subRouter.
				Path(sr.Pattern).
				Handler(handler).
				Methods(sr.Method).
				Name(sr.Name)
		}
	}

	return router
}

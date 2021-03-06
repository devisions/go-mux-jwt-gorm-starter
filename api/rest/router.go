package rest

import (
	"net/http"

	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/routes"
	"github.com/devisions/go-mux-jwt-gorm-starter/users"

	"github.com/gorilla/mux"
)

func NewRestApiRouter(usersRestApi users.RestApi) *mux.Router {

	router := mux.NewRouter()

	// setup the routes
	var ApiRestRouteSets []routes.ApiRestRouteSet

	ApiRestRouteSets = append(ApiRestRouteSets, usersRestApi.Routes)

	for _, routeSet := range ApiRestRouteSets {
		subRouter := router.PathPrefix(routeSet.Prefix).Subrouter()

		for _, sr := range routeSet.SubRoutes {
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

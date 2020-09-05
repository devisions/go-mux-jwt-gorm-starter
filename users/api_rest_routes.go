package users

import (
	"net/http"

	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/routes"
)

var ApiRestRouteSet = routes.ApiRestRouteSet{
	Prefix: "/users",
	SubRoutes: []routes.ApiRestRoute{
		{
			Name:        "Show All Users",
			Method:      http.MethodGet,
			Pattern:     "",
			HandlerFunc: ShowAllHandler,
			Protected:   true,
		},
		{
			Name:        "Show One User",
			Method:      http.MethodGet,
			Pattern:     "/{id}",
			HandlerFunc: ShowOneHandler,
			Protected:   true,
		},
		{
			Name:        "Register User",
			Method:      http.MethodPost,
			Pattern:     "",
			HandlerFunc: RegisterHandler,
			Protected:   false,
		},
		{
			Name:        "Update User",
			Method:      http.MethodPut,
			Pattern:     "",
			HandlerFunc: UpdateHandler,
			Protected:   false,
		},
		{
			Name:        "Delete User",
			Method:      http.MethodDelete,
			Pattern:     "{userId}",
			HandlerFunc: DeleteHandler,
			Protected:   true,
		},
		{
			Name:        "Login User",
			Method:      http.MethodPost,
			Pattern:     "/login",
			HandlerFunc: LoginHandler,
			Protected:   false,
		},
	},
}

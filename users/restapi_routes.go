package users

import (
	"net/http"

	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/routes"
)

func (handlers *RestApi) setupRestApiRoutes() {
	handlers.Routes = routes.ApiRestRouteSet{
		Prefix: "/users",
		SubRoutes: []routes.ApiRestRoute{
			{
				Name:        "Show All Users",
				Method:      http.MethodGet,
				Pattern:     "",
				HandlerFunc: handlers.ShowAllHandler,
				Protected:   true,
			},
			{
				Name:        "Show One User",
				Method:      http.MethodGet,
				Pattern:     "/{id}",
				HandlerFunc: handlers.showOne,
				Protected:   true,
			},
			{
				Name:        "Register User",
				Method:      http.MethodPost,
				Pattern:     "",
				HandlerFunc: handlers.RegisterHandler,
				Protected:   false,
			},
			{
				Name:        "Update User",
				Method:      http.MethodPut,
				Pattern:     "",
				HandlerFunc: handlers.UpdateHandler,
				Protected:   false,
			},
			{
				Name:        "Delete User",
				Method:      http.MethodDelete,
				Pattern:     "{userId}",
				HandlerFunc: handlers.DeleteHandler,
				Protected:   true,
			},
			{
				Name:        "Login User",
				Method:      http.MethodPost,
				Pattern:     "/login",
				HandlerFunc: handlers.LoginHandler,
				Protected:   false,
			},
		},
	}
}

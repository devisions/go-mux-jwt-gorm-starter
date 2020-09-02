package api

import (
	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/responses"
	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/routes"
	"github.com/devisions/go-mux-jwt-gorm-starter/user/domain"
	"net/http"
)

var UsersApiRestRoutes = routes.ApiRestRouteSet{
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
			Pattern:     "/{userId}",
			HandlerFunc: ShowOneHandler,
			Protected:   true,
		},
		{
			Name:        "CreateUser",
			Method:      http.MethodPost,
			Pattern:     "",
			HandlerFunc: CreateHandler,
			Protected:   false,
		},
		{
			Name:        "Update User",
			Method:      http.MethodPut,
			Pattern:     "",
			HandlerFunc: CreateHandler,
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
			Pattern:     "",
			HandlerFunc: LoginHandler,
			Protected:   false,
		},
	},
}

func ShowAllHandler(w http.ResponseWriter, r *http.Request) {
	var users []domain.User
	// repo.Find(&users)
	responses.RespondAsJson(w, users)
}

func ShowOneHandler(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var user domain.User
	// repo.Find(&user, params["id"]
	responses.RespondAsJson(w, user)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	user.Email = r.FormValue("email")
	user.Name = r.FormValue("name")
	user.Hash = user.HashPassword(r.FormValue("password"))
	// err := repo.Create(&user).Error
	// if err != nil {
	//    rest.RespondWithError(...
	// }
	responses.RespondAsJson(w, user)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// ...
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// ...
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// var user User
	// token, err := service.Login(email, password) ...

}

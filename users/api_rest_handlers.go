package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/responses"
	"github.com/devisions/go-mux-jwt-gorm-starter/app"
	"github.com/devisions/go-mux-jwt-gorm-starter/app/helpers"
	"github.com/gorilla/mux"
)

// UserService instance used internally (within)
var userSvc UserService

// InitApiRestHandlers does the initialization required for API REST Handlers to work correctly.
func InitApiRestHandlers(svc UserService) {
	userSvc = svc
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func ShowAllHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	// TODO
	responses.RespondJson(w, users)
}

func ShowOneHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.RespondErrBadRequest(w)
		return
	}
	user, err := userSvc.GetByID(uint(id))
	if err != nil {
		if err.Error() == string(app.ErrNotFound) {
			responses.RespondErrNotFound(w)
			return
		}
		responses.RespondJsonWithErrInternalAndReason(w, err.Error())
		return
	}
	responses.RespondJson(w, user)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := helpers.ParseForm(r, &form); err != nil {
		log.Printf("Error parsing signup request body: %s\n", err)
		responses.RespondJsonWithError(w, http.StatusBadRequest, "Invalid signup request")
		return
	}
	user, err := userSvc.Register(form.Name, form.Email, form.Password)
	if err != nil {
		responses.RespondJsonWithInternalServerError(w)
		return
	}
	responses.RespondJson(w, user)
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
	var form LoginForm
	if err := helpers.ParseForm(r, &form); err != nil {
		log.Printf("Error parsing login request body: %s\n", err)
		responses.RespondJsonWithError(w, http.StatusBadRequest, "Invalid login request")
		return
	}
	token, err := userSvc.Authenticate(form.Email, form.Password)
	if err != nil {
		log.Printf("Error at login: %s\n", err)
		responses.RespondJsonWithInternalServerError(w)
		return
	}
	responses.RespondJson(w, token)

}

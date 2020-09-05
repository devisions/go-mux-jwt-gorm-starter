package users

import (
	"log"
	"net/http"

	"github.com/devisions/go-mux-jwt-gorm-starter/api/helpers"
	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/responses"
)

// UserService instance used internally (within)
var userSvc UserService

// InitApiRestHandlers does the initialization required for API REST Handlers to work correctly.
func InitApiRestHandlers(svc UserService) {
	userSvc = svc
}

// SignupForm is submited when user does Signup.
type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func ShowAllHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	// TODO repo.Find(&users)
	responses.RespondJson(w, users)
}

func ShowOneHandler(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var user User
	// TODO repo.Find(&user, params["id"]
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

func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	// var user User
	// token, err := service.Login(email, password) ...

}

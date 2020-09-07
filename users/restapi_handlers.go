package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/responses"
	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/routes"
	"github.com/devisions/go-mux-jwt-gorm-starter/app"
	"github.com/devisions/go-mux-jwt-gorm-starter/app/helpers"
	"github.com/gorilla/mux"
)

type RestApi struct {
	Routes  routes.ApiRestRouteSet
	userSvc UserService
}

func NewRestApi(userSvc UserService) RestApi {
	restapi := RestApi{
		userSvc: userSvc,
	}
	restapi.setupRestApiRoutes()
	return restapi
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

func (h *RestApi) ShowAllHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	// TODO
	responses.RespondJson(w, users)
}

func (h *RestApi) showOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.RespondErrBadRequest(w)
		return
	}
	user, err := h.userSvc.GetByID(uint(id))
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

func (h *RestApi) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := helpers.ParseForm(r, &form); err != nil {
		log.Printf("Error parsing signup request body: %s\n", err)
		responses.RespondJsonWithError(w, http.StatusBadRequest, "Invalid signup request")
		return
	}
	user, err := h.userSvc.Register(form.Name, form.Email, form.Password)
	if err != nil {
		responses.RespondJsonWithInternalServerError(w)
		return
	}
	responses.RespondJson(w, user)
}

func (h *RestApi) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// ...
}

func (h *RestApi) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// ...
}

func (h *RestApi) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var form LoginForm
	if err := helpers.ParseForm(r, &form); err != nil {
		log.Printf("Error parsing login request body: %s\n", err)
		responses.RespondJsonWithError(w, http.StatusBadRequest, "Invalid login request")
		return
	}
	token, err := h.userSvc.Authenticate(form.Email, form.Password)
	if err != nil {
		log.Printf("Error at login: %s\n", err)
		responses.RespondJsonWithInternalServerError(w)
		return
	}
	responses.RespondJson(w, token)

}

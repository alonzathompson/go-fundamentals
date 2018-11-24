package userController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alonzathompson/go-fundamentals/mvcpattern/models"
	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	for i, v := range models.Users {
		fmt.Fprintf(w, "key: %s Value: %s \n", i, v)
	}

	/*w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n %s\n", i, v)*/
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Username: "james",
		Password: "jimdrty30",
		First:    "James",
		Last:     "Bond",
		ID:       "98786745632m",
	}

	uj, _ := json.Marshal(u)
	id := u.ID
	models.Users[id] = u

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n %s\n", uj, models.Users)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var data models.User
	//var users []models.User

	json.NewDecoder(r.Body).Decode(&data)

	models.Users[data.ID] = data

	fmt.Fprintf(w, "%s\n%s\n", data, models.Users)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var id = p.ByName("id")

	_, ok := models.Users[id]
	if ok {
		delete(models.Users, id)
	}

	fmt.Fprintf(w, "Make code work to delete a user %s\n %s\n", id, models.Users)
}

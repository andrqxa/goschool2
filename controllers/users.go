package controllers

import (
	"net/http"

	"goschool/models"
	"goschool/views"
)

// NewUser is used to create a new Users controller.
// This function will panic if the templates are not parsed correctly, and should
// be used during initial setup.
func NewUsers(us models.UserService) *Users {
	return &Users{
		NewView:   views.NewView("main", "users/new"),
		UsersView: views.NewView("main", "users/users"),
		us:        us,
	}
}

// Users ...
type Users struct {
	NewView   *views.View
	UsersView *views.View
	us        models.UserService
}

// New is used to render the form where a user can create a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	// type Alert struct {
	// 	Level   string
	// 	Message string
	// }
	// type Data struct {
	// 	Alert Alert
	// 	Yield interface{}
	// }
	// a := Alert{
	// 	Level:   "warning",
	// 	Message: "Successfully rendered a dynamic alert!",
	// }
	// d := Data{
	// 	Alert: a,
	// 	Yield: "hello!",
	// }
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

type SignupForm struct {
	Name             string
	Email            string
	RegistrationDate string
}

// Create is used to process the signup form when a user submits it. This is
// used to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name:             form.Name,
		Email:            form.Email,
		RegistrationDate: models.NewDate(),
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

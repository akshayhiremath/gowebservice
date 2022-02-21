package controllers

import (
	"net/http"
)

//Front controller
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	//Route beyond users/
	http.Handle("/users/", *uc)
}

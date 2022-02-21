package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	//Field to store a pattern to match incoming HTTP request path
	userIDPatttern *regexp.Regexp
}

//Method definition
//This method is bound to the userController struct using
//the uc userController parameter
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

//Constructor function. there are no actual constructors in Go
//Conventionally, in Go, contructor function names start with 'new' to indicate that its a constructor function
//Here the constructor is also returning a reference to the created object
//Go recognizes that an address of the variable is being returned and automatically
//promotes that variable up to the level it needs to be. So we are not gonna lose
//the track of userController variable created here
func newUserController() *userController {
	return &userController{
		userIDPatttern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

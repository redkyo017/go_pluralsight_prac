package controllers

import "net/http"

type UserController struct {
	Name string
}

func (uc *UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("con co be be"))
}

func NewUserController() *UserController {
	return &UserController{}
}

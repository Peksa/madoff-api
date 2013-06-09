package controllers

import "github.com/robfig/revel"
import "madoff-api/app/viewmodels"
import "encoding/json"

type Users struct {
	*revel.Controller
}

func (c Users) Add() revel.Result {
	// TODO: break out this type of code.
	var user viewmodels.UserViewModel
	json.NewDecoder(c.Request.Body).Decode(&user)
    return c.RenderJson(user)
}

func (c Users) List() revel.Result {
	var peksa, gunde, dschlyter viewmodels.UserViewModel

	peksa.Id = 1
	peksa.Username = "peksa"
	peksa.Name = "Niklas"

	gunde.Id = 2
	gunde.Username = "gunde"
	gunde.Name = "Henrik"

	dschlyter.Id = 3
	dschlyter.Username = "dschlyter"
	dschlyter.Name = "David"

	users := []viewmodels.UserViewModel{peksa, gunde, dschlyter}
	return c.RenderJson(users)
}

func (c Users) Show(id int) revel.Result {
	var user viewmodels.UserViewModel
	user.Id = id
	user.Username = "peksa"
	user.Name = "Niklas"
	return c.RenderJson(user)
}

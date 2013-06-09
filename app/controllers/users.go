package controllers

import "github.com/robfig/revel"
import "madoff-api/app/viewmodels"

type Users struct {
	*revel.Controller
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

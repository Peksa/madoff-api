package controllers

import "github.com/robfig/revel"
import "madoff-api/app/viewmodels"
import "madoff-api/app/models"
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

	peksa.Id = "1"
	peksa.Username = "peksa"
	peksa.Name = "Niklas"

	gunde.Id = "2"
	gunde.Username = "gunde"
	gunde.Name = "Henrik"

	dschlyter.Id = "3"
	dschlyter.Username = "dschlyter"
	dschlyter.Name = "David"

	users := []viewmodels.UserViewModel{peksa, gunde, dschlyter}
	return c.RenderJson(users)
}

func (c Users) Show(id string) revel.Result {
	// Magic to me.. Import random static method from models/user.go?
	user := models.Get(id); 

	var view viewmodels.UserViewModel
	view.Id = id
	view.Username = user.Username
	view.Name = user.Name
	str1 := "SEB"
	str2 := "111111111-1"
	str3 := "1234"
	view.BankName = &str1
	view.AccountNo = &str2
	view.ClearingNo = &str3
	return c.RenderJson(view)
}

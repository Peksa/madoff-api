package controllers

import "github.com/robfig/revel"
import "madoff-api/app/viewmodels"

type Receipts struct {
	*revel.Controller
}

func (c Receipts) List() revel.Result {
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

	var receipt1, receipt2, receipt3 viewmodels.ReceiptViewModel

	receipt1.Id = 1
	receipt1.Title = "Restaurant"
	receipt1.Description = "A nice restaurant"
	receipt1.Members = []viewmodels.UserViewModel{peksa, gunde}
	receipt1.Creator = gunde

	receipt2.Id = 2
	receipt2.Title = "Movies"
	receipt2.Description = "Great movie"
	receipt2.Members = []viewmodels.UserViewModel{dschlyter, gunde}
	receipt2.Creator = gunde

	receipt3.Id = 3
	receipt3.Title = "Bar"
	receipt3.Description = "Some bar"
	receipt3.Members = []viewmodels.UserViewModel{peksa, dschlyter, gunde}
	receipt3.Creator = dschlyter

	receipts := []viewmodels.ReceiptViewModel{receipt1, receipt2, receipt3}
	return c.RenderJson(receipts)
}

func (c Receipts) Show(id int) revel.Result {
	var peksa, dschlyter viewmodels.UserViewModel
	peksa.Id = 1
	peksa.Username = "peksa"
	peksa.Name = "Niklas"

	dschlyter.Id = 2
	dschlyter.Username = "dschlyter"
	dschlyter.Name = "David"

	var receipt viewmodels.ReceiptViewModel
	receipt.Id = id
	receipt.Title = "Movies"
	receipt.Description = "Great movie"
	receipt.Members = []viewmodels.UserViewModel{dschlyter, peksa}
	receipt.Creator = peksa
	return c.RenderJson(receipt)
}
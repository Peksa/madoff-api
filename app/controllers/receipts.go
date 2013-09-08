package controllers

import "github.com/robfig/revel"
import "madoff-api/app/viewmodels"

type Receipts struct {
	*revel.Controller
}

func (c Receipts) List() revel.Result {
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

	var receipt1, receipt2, receipt3 viewmodels.ReceiptViewModel

	receipt1.Id = "1"
	receipt1.Title = "Restaurant"
	desc1 := "A nice restaurant"
	receipt1.Description = &desc1
	var created1 int64 = 1378641336
	receipt1.Created = &created1
	receipt1.Members = &[]viewmodels.UserViewModel{peksa, gunde}
	//receipt1.Creator = gunde
	var comment1 viewmodels.CommentViewModel
	comment1.Id = "1"
	comment1.Content = "Yoyoyo nice!"
	comment1.Created = 1378651336
	comment1.Poster = peksa
	receipt1.Comments = &[]viewmodels.CommentViewModel{comment1}
	receipt1.Items = &[]viewmodels.ItemViewModel{}

	receipt2.Id = "2"
	receipt2.Title = "Great movie"
	desc2 := "A nice restaurant"
	receipt2.Description = &desc2
	var created2 int64 = 1378652121
	receipt2.Created = &created2
	receipt2.Members = &[]viewmodels.UserViewModel{dschlyter, gunde}
	receipt2.Creator = &gunde
	var comment2, comment3 viewmodels.CommentViewModel
	comment2.Id = "2"
	comment2.Content = "Wat is this?"
	comment2.Created = 1378652125
	comment2.Poster = gunde
	comment3.Id = "3"
	comment3.Content = "Det är nice!"
	comment3.Created = 1378652199
	comment3.Poster = dschlyter
	receipt2.Comments = &[]viewmodels.CommentViewModel{comment2, comment3}
	var item1, item2, item3 viewmodels.ItemViewModel
	item1.Id = "1"
	item1.Name = "3d-glasögon"
	item1.Amount = 25
	item1.Members = []viewmodels.UserViewModel{dschlyter}
	item2.Id = "2"
	item2.Name = "2d-glasögon"
	item2.Amount = 45
	item2.Members = []viewmodels.UserViewModel{gunde}
	item3.Id = "3"
	item3.Name = "Biobiljetter"
	item3.Amount = 220
	item3.Members = []viewmodels.UserViewModel{dschlyter, gunde}
	receipt2.Items = &[]viewmodels.ItemViewModel{item1, item2, item3}

	receipt3.Id = "3"
	receipt3.Title = "Bar"
	desc3 := "Some bar"
	receipt3.Description = &desc3
	var created3 int64 = 1368651235
	receipt2.Created = &created3
	receipt3.Members = &[]viewmodels.UserViewModel{peksa, dschlyter, gunde}
	receipt3.Creator = &dschlyter
	receipt3.Comments = &[]viewmodels.CommentViewModel{}
	var item4, item5 viewmodels.ItemViewModel
	item4.Id = "4"
	item4.Name = "3x125kr var"
	item4.Amount = 750
	item4.Members = []viewmodels.UserViewModel{peksa, gunde}
	item5.Id = "5"
	item5.Name = "3x120kr"
	item5.Amount = 360
	item5.Members = []viewmodels.UserViewModel{dschlyter}
	receipt3.Items = &[]viewmodels.ItemViewModel{item4, item5}

	receipts := []viewmodels.ReceiptViewModel{receipt1, receipt2, receipt3}
	return c.RenderJson(receipts)
}

func (c Receipts) Show(id string) revel.Result {
	var peksa, gunde viewmodels.UserViewModel
	peksa.Id = "1"
	peksa.Username = "peksa"
	peksa.Name = "Niklas"

	gunde.Id = "2"
	gunde.Username = "gunde"
	gunde.Name = "Henrik"

	var receipt1 viewmodels.ReceiptViewModel

	receipt1.Id = "1"
	receipt1.Title = "Restaurant"
	desc1 := "A nice restaurant"
	receipt1.Description = &desc1
	var created1 int64 = 1378641336
	receipt1.Created = &created1
	receipt1.Members = &[]viewmodels.UserViewModel{peksa, gunde}
	receipt1.Creator = &gunde
	var comment1 viewmodels.CommentViewModel
	comment1.Id = "1"
	comment1.Content = "Yoyoyo nice!"
	comment1.Created = 1378651336
	comment1.Poster = peksa
	receipt1.Comments = &[]viewmodels.CommentViewModel{comment1}
	receipt1.Items = &[]viewmodels.ItemViewModel{}
	return c.RenderJson(receipt1)
}
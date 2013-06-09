package viewmodels

type ReceiptViewModel struct {
	Id int						`json:"id"`
	Title string				`json:"title"`
	Description string			`json:"description"`
	Members []UserViewModel		`json:"members"`
	Creator UserViewModel		`json:"creator"`
}
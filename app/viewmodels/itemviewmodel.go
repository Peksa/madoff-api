package viewmodels

type ItemViewModel struct {
	// Peksa: ID för delete trots embedded?
	Id string                       `json:"id"`
	Name string                     `json:"name"`
	Amount float64                  `json:"amount"`
	Members []UserViewModel         `json:"members"`
}
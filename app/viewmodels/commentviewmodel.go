package viewmodels

type CommentViewModel struct {
	// Peksa: ID trots att de är embeddade? Tänker för potentiell delete...
	Id string                       `json:"id"`
	Created int64                   `json:"created"`
	Content string                  `json:"content"`
	Poster UserViewModel            `json:"poster"`
}
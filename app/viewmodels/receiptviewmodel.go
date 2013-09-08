package viewmodels

type ReceiptViewModel struct {
	Id string                       `json:"id"`
	Title string                    `json:"title"`
	
	// Optional  fields
	Description *string             `json:"description,omitempty"`
	// TODO peksa: kan man ha något date/time-objekt här och auto få
	// serialisering till unix timestamps?
	Created *int64                  `json:"created,omitempty"`
	Creator UserViewModel           `json:"creator,omitempty"`
	Members []UserViewModel         `json:"members,omitempty"`
	// Subrundor..
	Items []ItemViewModel           `json:"items,omitempty"`
	Comments []CommentViewModel     `json:"comments,omitempty"`

	// Calculated fields? Ska vi vara snälla och räkna ut lite saker
	// åt klienten? ish totalen av alla items etc.. 
}
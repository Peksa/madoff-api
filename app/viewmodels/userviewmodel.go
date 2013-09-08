package viewmodels

type UserViewModel struct {
	Id string               `json:"id"`
	Username string         `json:"username"`
	Name string             `json:"name"`
	// Optional
	BankName *string        `json:"bankName,omitempty"`
	AccountNo *string       `json:"accountNo,omitempty"`
	ClearingNo *string      `json:"clearingNo,omitempty"`
}
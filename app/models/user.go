package models

import "tux21b.org/v1/gocql"
import "tux21b.org/v1/gocql/uuid"

type User struct {
	Username string
	Name string
	Uuid uuid.UUID 
	Err error
}

func Get(id int) User {
	// Get a user from DB somehow..
	uuid := uuid.TimeUUID()
	err := gocql.ErrNotFound

	var ret User
	ret.Username = "peksa"
	ret.Name = "Peksa"
	ret.Uuid = uuid
	ret.Err = err
	return ret
}

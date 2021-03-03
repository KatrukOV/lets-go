package usr

import "log"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func (user *User) Validate() {
	log.Printf("User %v is validated", user)
}

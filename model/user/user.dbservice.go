package user

import (
	"errors"
	"fmt"
)

var userDB = make(map[int]*User)

// Save saves an user to data base
func (u *User) Save() error {
	if result := userDB[u.UserID]; result != nil {
		if result.Email == u.Email {
			return errors.New("User already exits")
		}
		return errors.New("User already exits")
	}
	userDB[u.UserID] = u
	fmt.Println(userDB[u.UserID])
	return nil
}

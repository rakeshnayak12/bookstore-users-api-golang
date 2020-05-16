package user

import (
	"github.com/rakeshnayak12/bookstore-users-api-golang/model/user"
)

// CreateUser contains the business logic to create am user
func CreateUser(u user.User) (*user.User, error) {
	if err := u.Save(); err != nil {
		return nil, err
	}
	return &u, nil
}

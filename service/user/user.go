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

// GetUser finds the result from the database
func GetUser(userID uint64) (*user.User, error) {
	result := &user.User{UserID: userID}
	if err := result.GetUser(); err != nil {
		return nil, err
	}
	return result, nil
}

package validation

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rakeshnayak12/bookstore-users-api-golang/model/user"
)

// CreateUser validates the request body for create user
func CreateUser(u *user.User, c *gin.Context) error {
	if err := c.ShouldBindJSON(&u); err != nil {
		return err
	}

	if u.Email == "" {
		return errors.New("Invalid email address")
	} else if u.FirstName == "" {
		return errors.New("First name required")
	} else if u.LastName == "" {
		return errors.New("Last name required")
	}
	return nil
}
package validation

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rakeshnayak12/bookstore-users-api-golang/model/user"
)

// CreateUser validates the request body for create user route
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

// GetUser validates the request body for getuser route
func GetUser(c *gin.Context) (uint64, error) {
	userID, err := strconv.ParseUint(c.Param("user-id"), 10, 64)
	if err != nil || userID == 0 {
		return 0, errors.New("userId should be a positive number")
	}
	return userID, nil
}

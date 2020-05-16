package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/rakeshnayak12/bookstore-users-api-golang/model/user"
	userService "github.com/rakeshnayak12/bookstore-users-api-golang/service/user"
	customError "github.com/rakeshnayak12/bookstore-users-api-golang/utils/error"
	validate "github.com/rakeshnayak12/bookstore-users-api-golang/validation"
)

// GetUsers finds all the users
func GetUsers(c *gin.Context) {
	var user userModel.User
	fmt.Println(user)

}

// CreateUser creates a user
func CreateUser(c *gin.Context) {
	var user userModel.User

	if err := validate.CreateUser(&user, c); err != nil {
		restErr := customError.HandleError("bad request", err.Error(), http.StatusBadRequest)
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	result, err := userService.CreateUser(user)
	if err != nil {
		restErr := customError.HandleError("bad request", err.Error(), http.StatusInternalServerError)
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetUser finds a specific user
func GetUser(c *gin.Context) {

}

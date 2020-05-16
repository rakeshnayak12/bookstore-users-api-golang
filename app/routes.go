package app

import (
	"github.com/rakeshnayak12/bookstore-users-api-golang/controller/user"
)

func routes() {
	router.GET("/users/get-users", user.GetUsers)
	router.GET("/users/get-user/:userid", user.GetUser)
	router.POST("/users/create-user", user.CreateUser)
}

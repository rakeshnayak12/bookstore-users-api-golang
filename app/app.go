package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp starts the application
func StartApp() {
	routes()
	router.Run(":8081")
}

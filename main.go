package main

import (
	entities "github.com/Spid3rrr/basic-backend-in-go/entities"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	entities.SetupEntityRoutes(router)
	router.Run("localhost:8080")
}

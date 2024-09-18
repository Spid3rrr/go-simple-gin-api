package entities

import (
	galaxies "github.com/Spid3rrr/basic-backend-in-go/entities/galaxies"
	planets "github.com/Spid3rrr/basic-backend-in-go/entities/planets"
	"github.com/gin-gonic/gin"
)

func SetupEntityRoutes(router *gin.Engine) {
	planets.RegisterPlanetRoutes(router)
	galaxies.RegisterGalaxyRoutes(router)
}

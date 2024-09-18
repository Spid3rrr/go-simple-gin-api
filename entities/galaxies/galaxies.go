package galaxies

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type Galaxy struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Distance int    `json:"distance" validate:"required"`
}

var galaxies = []Galaxy{
	Galaxy{ID: "1", Name: "Milky Way", Distance: 1000},
	Galaxy{ID: "2", Name: "Andromeda", Distance: 2000},
	Galaxy{ID: "3", Name: "Triangulum", Distance: 3000},
}

func getGalaxies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, galaxies)
}

func getGalaxyById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range galaxies {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "galaxy not found"})
}

func addGalaxy(c *gin.Context) {
	var newGalaxy Galaxy

	if err := c.BindJSON(&newGalaxy); err != nil {
		return
	}

	// Validate newGalaxy fields
	if err := validate.Struct(newGalaxy); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "validation failed", "errors": err.Error()})
		return
	}

	galaxies = append(galaxies, newGalaxy)
	c.IndentedJSON(http.StatusCreated, newGalaxy)
}

func deleteGalaxy(c *gin.Context) {
	id := c.Param("id")
	for i, a := range galaxies {
		if a.ID == id {
			galaxies = append(galaxies[:i], galaxies[i+1:]...)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "galaxy not found"})
}

func RegisterGalaxyRoutes(router *gin.Engine) {
	validate = validator.New()
	galaxiesGroup := router.Group("/galaxies")
	{
		galaxiesGroup.GET("/", getGalaxies)
		galaxiesGroup.GET("/:id", getGalaxyById)
		galaxiesGroup.POST("/", addGalaxy)
		galaxiesGroup.DELETE("/:id", deleteGalaxy)
	}
}

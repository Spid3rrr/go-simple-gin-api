package planets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type Planet struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Distance int    `json:"distance" validate:"required"`
}

var planets = []Planet{
	Planet{ID: "1", Name: "Mercury", Distance: 36},
	Planet{ID: "2", Name: "Venus", Distance: 67},
	Planet{ID: "3", Name: "Earth", Distance: 93},
}

func getPlanets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, planets)
}

func getPlanetById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range planets {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "planet not found"})
}

func addPlanet(c *gin.Context) {
	var newPlanet Planet

	if err := c.BindJSON(&newPlanet); err != nil {
		return
	}

	// Validate newPlanet fields
	if err := validate.Struct(newPlanet); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "validation failed", "errors": err.Error()})
		return
	}

	planets = append(planets, newPlanet)
	c.IndentedJSON(http.StatusCreated, newPlanet)
}

func deletePlanet(c *gin.Context) {
	id := c.Param("id")
	for i, a := range planets {
		if a.ID == id {
			planets = append(planets[:i], planets[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "planet deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "planet not found"})
}

func RegisterPlanetRoutes(router *gin.Engine) {
	validate = validator.New()

	planetRoutes := router.Group("/planets")
	{
		planetRoutes.GET("", getPlanets)
		planetRoutes.GET("/:id", getPlanetById)
		planetRoutes.POST("", addPlanet)
		planetRoutes.DELETE("/:id", deletePlanet)
	}
}

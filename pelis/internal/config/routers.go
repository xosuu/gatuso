package config

import (
	"pelis/internal/controler"

	"github.com/gin-gonic/gin"
)

func LoadRouters(controller *controler.MovieController)(*gin.Engine ){

	routers := gin.Default() 
	api := routers.Group("/api")
	
	api.GET("/movies", controller.GetAllMovies)
	api.GET("/movies/:id", controller.GetById)
	api.POST("/movies", controller.InsertMovie)
	api.DELETE("movies/:id", controller.DeleteById)
	api.PUT("/movies", controller.UpdateMovie)

	return routers
}
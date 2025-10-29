package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/rquiogue/travel-to-do-list/internal/controllers"
	"github.com/rquiogue/travel-to-do-list/internal/repositories"
)

func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	repo := repositories.NewLocationRepository(db)
	controller := controllers.NewLocationController(repo)

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000", "https://yourfrontend.com"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        AllowCredentials: true,
    }))

	r.GET("/locations", controller.GetLocations)
	r.POST("/locations", controller.CreateLocation)
	r.PUT("/locations/:id", controller.UpdateLocation)
	r.DELETE("/locations/:id", controller.DeleteLocation)

	return r
}

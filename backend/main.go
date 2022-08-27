package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"straccia17.com/box-catalog-api/controllers"
	"straccia17.com/box-catalog-api/services"
)

func isLocalDevelopment() bool {
	env := os.Getenv("APP_ENV")
	return env == ""
}

func main() {

	localDev := isLocalDevelopment()

	if localDev {
		log.Println("Load local environment variables")
		godotenv.Load(".env.local")
	}

	services.InitDB()

	router := gin.Default()

	if localDev {

		log.Println("Configure CORS middleware")

		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"http://box-catalog.straccia17.com:5173"}
		config.AllowCredentials = true

		router.Use(cors.New(config))
	}

	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	api := router.Group("/").Use(services.VerifyJWT())

	api.GET("/locations", controllers.GetLocations)
	api.POST("/locations", controllers.NewLocation)

	api.GET("/categories", controllers.GetCategories)
	api.POST("/categories", controllers.NewCategory)
	api.GET("/categories/:categoryId/items", controllers.GetCategoryItems)

	api.GET("/boxes", controllers.GetBoxes)
	api.POST("/boxes", controllers.NewBox)
	api.GET("/boxes/:boxId/items", controllers.GetBoxItems)

	api.GET("/items", controllers.GetItems)
	api.POST("/items", controllers.NewItem)

	router.Run(":8080")
}

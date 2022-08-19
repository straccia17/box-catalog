package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"straccia17.com/box-catalog-api/controllers"
	"straccia17.com/box-catalog-api/services"
)

func initEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	if env != "test" {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load()
}

func main() {
	initEnv()
	services.InitDB()

	router := gin.Default()

	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	api := router.Group("/").Use(services.VerifyJWT())

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

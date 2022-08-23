package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"straccia17.com/box-catalog-api/models"
	"straccia17.com/box-catalog-api/services"
)

func GetLocations(c *gin.Context) {
	userId, _ := services.RetrieveUserInfo(c)
	locations := make([]models.Location, 0)
	rows, err := services.DB.Queryx("SELECT location_id, label FROM locations WHERE user_id = $1", userId)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve locations"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Location
		err := rows.StructScan(&b)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve locations"})
		}
		locations = append(locations, b)
	}
	c.IndentedJSON(http.StatusOK, locations)
}

func NewLocation(c *gin.Context) {
	var newLocation models.NewLocation

	if err := c.ShouldBindJSON(&newLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, _ := services.RetrieveUserInfo(c)
	const newLocationStmt = "INSERT INTO locations(label, user_id) VALUES ($1, $2)"

	_, err := services.DB.Exec(newLocationStmt, newLocation.Label, userId)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to save box"})
		return
	}

	c.Status(http.StatusCreated)
}

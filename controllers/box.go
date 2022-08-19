package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"straccia17.com/box-catalog-api/models"
	"straccia17.com/box-catalog-api/services"
)

func GetBoxes(c *gin.Context) {
	userId, _ := services.RetrieveUserInfo(c)
	boxes := make([]models.Box, 0)
	rows, err := services.DB.Queryx("SELECT box_id, label, position FROM boxes WHERE user_id = $1", userId)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve boxes"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Box
		err := rows.StructScan(&b)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve boxes"})
		}
		boxes = append(boxes, b)
	}
	c.IndentedJSON(http.StatusOK, boxes)
}

func NewBox(c *gin.Context) {
	var newBox models.NewBox

	if err := c.ShouldBindJSON(&newBox); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, _ := services.RetrieveUserInfo(c)
	const newBoxStmt = "INSERT INTO boxes(label, position, user_id) VALUES ($1, $2, $3)"

	_, err := services.DB.Exec(newBoxStmt, newBox.Label, newBox.Position, userId)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to save box"})
		return
	}

	c.Status(http.StatusCreated)
}

func GetBoxItems(c *gin.Context) {
	userId, _ := services.RetrieveUserInfo(c)
	items := make([]models.Item, 0)
	boxId := c.Param("boxId")
	rows, err := services.DB.Queryx(`
	SELECT
		i.item_id,
		i.item,
		c.title as category,
		b.label as box,
		b.position 
	FROM items i 
	LEFT JOIN boxes b ON i.box_id = b.box_id
	LEFT JOIN categories c ON i.category_id = c.category_id
	WHERE
		i.box_id = $1
		AND i.user_id = $2
	`, boxId, userId)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve items"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var it models.Item
		err := rows.StructScan(&it)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve items"})
		}
		items = append(items, it)
	}
	c.IndentedJSON(http.StatusOK, items)
}

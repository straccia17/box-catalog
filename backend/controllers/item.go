package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"straccia17.com/box-catalog-api/models"
	"straccia17.com/box-catalog-api/services"
)

func GetItems(c *gin.Context) {
	userId, _ := services.RetrieveUserInfo(c)
	items := make([]models.Item, 0)
	rows, err := services.DB.Queryx(`
	SELECT
		i.item_id,
		i.item,
		c.title as category,
		b.label as box,
		l.label as location
	FROM items i 
	LEFT JOIN boxes b ON i.box_id = b.box_id
	LEFT JOIN locations l ON b.location_id = l.location_id
	LEFT JOIN categories c ON i.category_id = c.category_id
	WHERE i.user_id = $1
	`, userId)
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

func NewItem(c *gin.Context) {
	var newItem models.NewItem

	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, _ := services.RetrieveUserInfo(c)
	const newItemStmt = "INSERT INTO items(item, box_id, category_id, user_id) VALUES ($1, $2, $3, $4)"

	_, err := services.DB.Exec(newItemStmt, newItem.Item, newItem.BoxID, newItem.CategoryID, userId)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to save item"})
		return
	}

	c.Status(http.StatusCreated)
}

package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"straccia17.com/box-catalog-api/models"
	"straccia17.com/box-catalog-api/services"
)

func GetCategories(c *gin.Context) {
	userId, _ := services.RetrieveUserInfo(c)
	categories := make([]models.Category, 0)
	rows, err := services.DB.Queryx("SELECT category_id, title FROM categories WHERE user_id = $1", userId)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve categories"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cat models.Category
		err := rows.StructScan(&cat)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve categories"})
		}
		categories = append(categories, cat)
	}
	c.IndentedJSON(http.StatusOK, categories)
}

func NewCategory(c *gin.Context) {
	var newCategory models.NewCategory

	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, _ := services.RetrieveUserInfo(c)
	const newCategoryStmt = "INSERT INTO categories(title, user_id) VALUES ($1, $2)"

	_, err := services.DB.Exec(newCategoryStmt, newCategory.Title, userId)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to save category"})
		return
	}

	c.Status(http.StatusCreated)
}

func GetCategoryItems(c *gin.Context) {
	userId, _ := services.RetrieveUserInfo(c)
	items := make([]models.Item, 0)
	categoryId := c.Param("categoryId")
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
		i.category_id = $1
		AND i.user_id = $2
	`, categoryId, userId)
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

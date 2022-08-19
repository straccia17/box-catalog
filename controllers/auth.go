package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"straccia17.com/box-catalog-api/services"
)

type Credential struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID       string `json:"id" db:"user_id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"-" db:"password"`
}

func RegisterUser(c *gin.Context) {
	var input Credential

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := services.DB.QueryRowx(`
	SELECT
		u.email
	FROM users u 
	WHERE
		u.email = $1
	`, input.Email)

	var email string
	err := row.Scan(&email)

	if err == sql.ErrNoRows {
		const newUserStmt = "INSERT INTO users VALUES (gen_random_uuid(), $1, $2)"
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
		_, err := services.DB.Exec(newUserStmt, input.Email, hashedPassword)

		if err != nil {
			log.Println(err.Error())
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to register new user"})
			return
		}

		c.Status(http.StatusCreated)

	} else if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to retrieve users"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "email already used"})
	}
}

func LoginUser(c *gin.Context) {
	var input Credential

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := services.DB.QueryRowx(`
	SELECT
		*
	FROM users u 
	WHERE
		u.email = $1
	`, input.Email)

	var u User
	err := row.StructScan(&u)

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})

	} else if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
	} else {

		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password))
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
			return
		}

		token, err := services.GenerateJWT(u.ID, u.Email)

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"token": token})
	}
}

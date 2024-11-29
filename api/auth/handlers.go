package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/database"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/models"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var userInput models.UserInput
	c.BindJSON(&userInput)

	fmt.Println(userInput)

	user := models.User{
		Username:  userInput.Username,
		Password:  userInput.Password,
		FirstName: userInput.FirstName,
		LastName:  userInput.LastName,
		Email:     userInput.Email,
	}
	dbConn, _ := database.CreateDbConnection()
	if tx := dbConn.Create(&user); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{
		"username": user.Username,
		"id":       user.ID,
	})
}

func Login(c *gin.Context) {
	var input models.UserInput
	var user models.User

	c.BindJSON(&input)
	dbConn, _ := database.CreateDbConnection()
	dbConn.Where("username=?", input.Username).Find(&user)

	fmt.Println("USUARIO ENCONTRADO: ", user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error Invalid credentials": err.Error()})
		return
	}

	sessionToken := uuid.NewV5(uuid.UUID{}, user.Username).String()

	session := models.Session{
		Uid:        user.ID,
		ExpiryTime: time.Now().Add(10 * time.Minute),
	}

	Sessions[sessionToken] = session

	claims := models.Payload{
		MapClaims: jwt.MapClaims{
			"iat": jwt.NewNumericDate(time.Now()),
			"eat": jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			"id":  user.ID,
		},
		Session:       sessionToken,
		SearchFilters: user.SearchFilters,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signinKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	tokenString, err := token.SignedString(signinKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

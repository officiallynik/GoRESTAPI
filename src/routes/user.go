package routes

import (
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// User ==> user model
type User struct {
	ID       string
	Email    string
	Password string
}

// Users ==> users in-memory store
var Users = make(map[string]User)

// RegisterUser ==> logic to register users
func RegisterUser(c *gin.Context) {
	body := c.Request.Body

	var user User
	json.NewDecoder(body).Decode(&user)

	if _, ok := Users[user.Email]; ok {
		c.JSON(400, gin.H{
			"data": "Account with given email id already exists",
		})

	} else {
		UUIDUser := uuid.New()
		user.ID = strings.ReplaceAll(UUIDUser.String(), "-", "")

		pwdhash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

		if err != nil {
			fmt.Println("Problem creating user")
			c.JSON(400, gin.H{
				"data": "Failed to register user",
			})
			return
		}
		user.Password = string(pwdhash)

		fmt.Println(user)

		Users[user.Email] = user

		c.JSON(201, gin.H{
			"data": "account created successfully",
		})
	}
}

// LoginUser ==> logic to login user
func LoginUser(c *gin.Context) {
	body := c.Request.Body

	var user User
	json.NewDecoder(body).Decode(&user)

	email := user.Email
	pwdreal := Users[email].Password
	err := bcrypt.CompareHashAndPassword([]byte(pwdreal), []byte(user.Password))

	if err != nil {
		// fmt.Println("authentication failed")
		c.JSON(400, gin.H{
			"data": "authentication failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": Users[email].ID,
	})
}

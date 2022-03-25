package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

func GetJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "ok"})
}

func GetUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, User{ID: 1, Name: "Alex", Username: "alexbt"})
}

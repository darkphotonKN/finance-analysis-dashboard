package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{"message": "User created test."})
}

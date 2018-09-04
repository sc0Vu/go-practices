package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Pong"})
}

// Health
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": true})
}
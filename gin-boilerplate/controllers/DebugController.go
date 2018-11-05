package controllers

import (
	"net/http"

	"gin-boilerplate/types"

	"github.com/gin-gonic/gin"
)

// Ping debug function for ping
func Ping(c *gin.Context) {
	response := types.APIResponse{Msg: "Pong", Success: true}
	c.JSON(http.StatusOK, response)
}

// Health debug function for health
func Health(c *gin.Context) {
	response := types.APIResponse{Msg: true, Success: true}
	c.JSON(http.StatusOK, response)
}

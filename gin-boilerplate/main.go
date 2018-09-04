package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"gin-boilerplate/config"
	"gin-boilerplate/controllers"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Set session
	store := sessions.NewCookieStore([]byte("session"))
	r.Use(sessions.Sessions("session", store))

	// Setup cors
	// Use this when in production
	corsConfig := cors.Config{
		AllowOrigins:     config.API.CORSDomains,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Set mode
	if config.API.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	} else {
		debugRouter := r.Group("debug")
		{
			debugRouter.GET("/ping", controllers.Ping)
			debugRouter.GET("/health", controllers.Health)
		}
	}
	r.Use(cors.New(corsConfig))

	// Listen and Server in config.API.Domain:config.API.Port
	r.Run(config.API.Domain + ":" + config.API.Port)
}

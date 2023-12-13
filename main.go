// cmd/main.go

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"logger/app/utils"
	"logger/app/routers"
)

func main() {
	// Initialize configuration
	cfg := utils.NewConfigPort()
	cfg.LoadFromEnv()

	// Set up Gin router
	router := gin.New()

	// Forward routes to services
	routers.SetupRoutes(router) // Include QR routes

	// Run the server
	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
